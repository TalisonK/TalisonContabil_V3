package br.com.talison.contabil.service;

import br.com.talison.contabil.domain.Expense;
import br.com.talison.contabil.domain.Income;
import br.com.talison.contabil.domain.Totals;
import br.com.talison.contabil.domain.dto.ActivityDto;
import br.com.talison.contabil.domain.dto.DashboardDto;
import br.com.talison.contabil.domain.dto.TotalsDto;
import br.com.talison.contabil.repository.ExpenseRepository;
import br.com.talison.contabil.repository.IncomeRepository;
import br.com.talison.contabil.repository.TotalsRepository;
import br.com.talison.contabil.service.mapper.ActivityExpenseMapper;
import br.com.talison.contabil.service.mapper.ActivityIncomeMapper;
import br.com.talison.contabil.service.mapper.TotalsMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.text.DecimalFormat;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.*;

@Service
@Transactional
@RequiredArgsConstructor
public class TotalsService {

    private final TotalsRepository totalsRepository;
    private final IncomeRepository incomeRepository;
    private final ExpenseRepository expenseRepository;

    private final ActivityIncomeMapper activityIncomeMapper;
    private final ActivityExpenseMapper activityExpenseMapper;
    private final TotalsMapper totalsMapper;

    private List<String> months = Arrays.asList("Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep","Oct","Nov","Dec");

    private Integer convertMonthToNumber(String month) {

        int selectedMonth = 0;

        if(months.contains(month)) {
            selectedMonth = months.indexOf(month) + 1;
        }
        else {
            Dictionary<String, String> monthsDic = months();

            String translate = monthsDic.get(month);

            selectedMonth = months.indexOf(translate) + 1;
        }

        return selectedMonth;

    }

    private Dictionary<String, String> months(){
        Dictionary<String, String> months = new Hashtable<>();
        months.put("Fev", "Feb");
        months.put("Abr", "Apr");
        months.put("Mai", "May");
        months.put("Ago", "Aug");
        months.put("Set", "Sep");
        months.put("Out", "Oct");
        months.put("Dez", "Dec");
        return months;
    }

    private String convertMonthToString(Integer month) {
        return months.get(month - 1);
    }

    private List<Date> calendarGenerator(String year, String month) {

        String start = "";
        String end = "";

        if(month.length() == 3){
            start = year + "-" + convertMonthToNumber(month) + "-01";
            end = year + "-" + convertMonthToNumber(month) + "-31";
        }
        else {
            start = year + "-" + month + "-01";
            end = year + "-" + month + "-31";
        }

        SimpleDateFormat form = null;

        try{
            form = new SimpleDateFormat("yyyy-MM-dd");
            Date Date1 = form.parse(start);
            Date Date2 = form.parse(end);
            return Arrays.asList(Date1, Date2);
        }
        catch (ParseException e){
            System.out.println("Erro ao gerar as datas de filtragem! Erro:" + year + month);;
        }
        return null;
    }

    private Totals createIncome(String year, String month, String userId){

        Optional<Totals> income = totalsRepository.findByTypeAndYearAndMonthAndUserId("income", year, month, userId);

        if(income.isPresent()){
            return income.get();
        }
        else {
            Totals totals = new Totals();
            totals.setYear(year);
            totals.setMonth(month);
            totals.setUpdatedAt(new Date());
            totals.setCreatedAt(new Date());
            totals.setUserId(userId);

            List<Date> calendar = calendarGenerator(year, month);

            Optional<List<Income>> value = incomeRepository.findAllByUserIdAndReceivedAtBetweenOrderByReceivedAt(userId, calendar.get(0), calendar.get(1));

            Double cont = value.get().stream().map(Income::getValue).reduce(0.0, Double::sum);

            DecimalFormat df = new DecimalFormat("####0.00");

            totals.setType("income");
            totals.setValue(Double.parseDouble(df.format(cont).replace(",", ".")));
            totalsRepository.save(totals);
            return totals;
        }
    }

    private Totals createExpense(String year, String month, String userId){

        Optional<Totals> expense = totalsRepository.findByTypeAndYearAndMonthAndUserId("expense", year, month, userId);

        if(expense.isPresent()){
            return expense.get();
        }
        else {
            Totals totals = new Totals();
            totals.setYear(year);
            totals.setMonth(month);
            totals.setUpdatedAt(new Date());
            totals.setCreatedAt(new Date());
            totals.setUserId(userId);

            List<Date> calendar = calendarGenerator(year, month);

            Optional<List<Expense>> value = expenseRepository.findAllByUserIdAndPaidAtBetweenOrderByPaidAt(userId, calendar.get(0), calendar.get(1));

            Double cont = value.get().stream().map(Expense::getValue).reduce(0.0, Double::sum);


            DecimalFormat df = new DecimalFormat("####0.00");

            totals.setType("expense");
            totals.setValue(Double.parseDouble(df.format(cont).replace(",", ".")));
            totalsRepository.save(totals);
            return totals;
        }
    }

    public DashboardDto dashboardHandler(String userId, String year, String month){

        DashboardDto dashboardDto = new DashboardDto(userId, year, month);

        dashboardDto.getResumes().put("incomes", getLastIncomeTotals(userId, year, month));
        dashboardDto.getResumes().put("expenses", getLastExpenseTotals(userId, year, month));
        dashboardDto.getResumes().put("balances", getLastBalanceTotals(userId, year, month));
        dashboardDto.getTimeline().addAll(timelineByMonth(userId, year, month));

        dashboardDto.getCharts().put("incomeXexpense", getIncomeVsExpense(userId, year, month));

        return dashboardDto;
    }

    public TotalsDto getTotals(String year, String month, String userId, String type){

        Dictionary<String, String> months = months();

        String mon = months.get(month) == null ? month : months.get(month);

        if(type.equals("income")){
            return totalsMapper.toDto(createIncome(year, mon, userId));
        }
        else{
            return totalsMapper.toDto(createExpense(year, mon, userId));
        }
    }

    public TotalsDto updateTotals(LocalDateTime date, String userId, String type){

        Date converter = Date.from(date.atZone(ZoneId.systemDefault()).toInstant());

        String month = new SimpleDateFormat("MMM", Locale.ENGLISH).format(converter);
        String year = new SimpleDateFormat("yyyy", Locale.ENGLISH).format(converter);

        return updateTotals(year, month, userId, type);
    }


    public TotalsDto updateTotals(String year, String month, String userId, String type){

        Optional<Totals> explorer = totalsRepository.findByTypeAndYearAndMonthAndUserId(type, year, month, userId);

        if(explorer.isEmpty()){
            return getTotals(year, month, userId, type);
        }
        else{
            Totals totals = explorer.get();
            totals.setUpdatedAt(new Date());

            if(Objects.equals(type, "income")){
                List<Date> calendar = calendarGenerator(year, month);

                Optional<List<Income>> value = incomeRepository.findAllByUserIdAndReceivedAtBetweenOrderByReceivedAt(userId, calendar.get(0), calendar.get(1));
                Double cont = value.get().stream().map(Income::getValue).reduce(0.0, Double::sum);

                totals.setValue(cont);

                totalsRepository.save(totals);

            }
            else if(Objects.equals(type, "expense")){
                List<Date> calendar = calendarGenerator(year, month);

                Optional<List<Expense>> value = expenseRepository.findAllByUserIdAndPaidAtBetweenOrderByPaidAt(userId, calendar.get(0), calendar.get(1));
                Double cont = value.get().stream().map(Expense::getValue).reduce(0.0, Double::sum);

                totals.setValue(cont);

                totalsRepository.save(totals);
            }

            return totalsMapper.toDto(totals);
        }
    }

    public List<TotalsDto> getLastIncomeTotals(String userId, String year, String month){
        return getLastTotals(userId, year, month, "income");
    }

    public List<TotalsDto> getLastExpenseTotals(String userId, String year, String month){
        return getLastTotals(userId, year, month, "expense");
    }

    public List<TotalsDto> getLastTotals(String userId, String year, String month, String type){
        TotalsDto actual = getTotals(year, month, userId, type);

        List<TotalsDto> totals = new ArrayList<>();

        for(int i = 1; i > 0; i--){
            int monthAux = convertMonthToNumber(month) - i;
            String yearAux = year;

            if(monthAux < 1){
                monthAux = 12 + monthAux;
                yearAux = String.valueOf(Integer.parseInt(yearAux) - 1);
            }
            else if(monthAux > 12){
                monthAux = monthAux - 12;
                yearAux = String.valueOf(Integer.parseInt(yearAux) + 1);
            }

            totals.add(getTotals(yearAux, convertMonthToString(monthAux), userId, type));
        }

        totals.add(actual);
        return totals;
    }

    public List<TotalsDto> getLastBalanceTotals(String userId, String year, String month) {

        List<TotalsDto> incomes = getLastIncomeTotals(userId, year, month);
        List<TotalsDto> expenses = getLastExpenseTotals(userId, year, month);

        List<TotalsDto> balances = new ArrayList<>();

        for (int i = 0; i < incomes.size(); i++) {
            TotalsDto balance = new TotalsDto();
            balance.setYear(incomes.get(i).getYear());
            balance.setMonth(incomes.get(i).getMonth());
            balance.setUserId(incomes.get(i).getUserId());
            balance.setValue(incomes.get(i).getValue() - expenses.get(i).getValue());
            balance.setType("balance");
            balances.add(balance);
        }

        return balances;
    }

    public HashMap<String, List<TotalsDto>> getIncomeVsExpense(String userId, String year, String month){

        HashMap<String, List<TotalsDto>> incomeVsExpense = new HashMap<>();

        List<TotalsDto> incomes = new ArrayList<>();
        List<TotalsDto> expenses = new ArrayList<>();

        int monthAux = convertMonthToNumber(month);

        String yearAux = year;

        for(int i = 0; i < 7; i++){
            if(monthAux < 1){
                monthAux = 12;
                yearAux = String.valueOf(Integer.parseInt(yearAux) - 1);
            }
            incomes.add(getTotals(yearAux, convertMonthToString(monthAux), userId, "income"));
            expenses.add(getTotals(yearAux, convertMonthToString(monthAux), userId, "expense"));
            monthAux--;
        }

        monthAux = convertMonthToNumber(month) + 1;
        yearAux = year;

        Collections.reverse(incomes);
        Collections.reverse(expenses);

        for (int i = 0; i < 6; i++) {
            if (monthAux > 12) {
                monthAux = 1;
                yearAux = String.valueOf(Integer.parseInt(yearAux) + 1);
            }
            incomes.add(getTotals(yearAux, convertMonthToString(monthAux), userId, "income"));
            expenses.add(getTotals(yearAux, convertMonthToString(monthAux), userId, "expense"));
            monthAux++;
        }

        incomeVsExpense.put("incomes", incomes);
        incomeVsExpense.put("expenses", expenses);

        return incomeVsExpense;

    }

    public List<ActivityDto> timelineByMonth(String userId, String year, String month) {

        List<Date> calendar = calendarGenerator(year, month);

        List<Income> incomes = incomeRepository.findAllByUserIdAndReceivedAtBetweenOrderByReceivedAt(userId, calendar.get(0), calendar.get(1)).orElse(new ArrayList<>());
        List<Expense> expenses = expenseRepository.findAllByUserIdAndPaidAtBetweenOrderByPaidAt(userId, calendar.get(0), calendar.get(1)).orElse(new ArrayList<>());


        List<ActivityDto> activities = new ArrayList<>();
        int inc = 0;
        int exp = 0;

        while(inc < incomes.size() || exp < expenses.size()){

            ActivityDto activity;
            if(inc == incomes.size()){
                activity = activityExpenseMapper.toDto(expenses.get(exp));
                activity.setType("expense");
                exp++;
            }
            else if (exp == expenses.size()) {
                activity = activityIncomeMapper.toDto(incomes.get(inc));
                activity.setType("income");
                inc++;
            }
            else{
                if(incomes.get(inc).getReceivedAt().compareTo(expenses.get(exp).getPaidAt()) < 0){
                    activity = activityIncomeMapper.toDto(incomes.get(inc));
                    activity.setType("income");
                    inc++;
                }
                else{
                    activity = activityExpenseMapper.toDto(expenses.get(exp));
                    activity.setType("expense");
                    exp++;
                }
            }

            activities.add(activity);
        }
        activities.sort(Comparator.comparing(ActivityDto::getCreatedAt).reversed());
        return activities;
    }
}
