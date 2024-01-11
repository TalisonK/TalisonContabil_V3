package br.com.talison.contabil.service;

import br.com.talison.contabil.domain.Expense;
import br.com.talison.contabil.domain.Income;
import br.com.talison.contabil.domain.Totals;
import br.com.talison.contabil.domain.dto.*;
import br.com.talison.contabil.repository.ExpenseRepository;
import br.com.talison.contabil.repository.IncomeRepository;
import br.com.talison.contabil.repository.TotalsRepository;
import br.com.talison.contabil.service.mapper.ActivityExpenseMapper;
import br.com.talison.contabil.service.mapper.ActivityIncomeMapper;
import br.com.talison.contabil.service.mapper.ExpenseMapper;
import br.com.talison.contabil.service.mapper.TotalsMapper;
import br.com.talison.contabil.service.utils.DateUtils;
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
    private final ExpenseMapper expenseMapper;
    private final TotalsMapper totalsMapper;

    private final DateUtils dateUtils = new DateUtils();

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

            List<Date> calendar = dateUtils.calendarGenerator(year, month);

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

            List<Date> calendar = dateUtils.calendarGenerator(year, month);

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

        dashboardDto.getIncomeVSexpense().addAll(getIncomeVsExpense(userId, year, month));
        dashboardDto.getExpenseByCategory().putAll(getExpenseByCategory(userId, year, month, dashboardDto.getTimeline()));
        dashboardDto.getExpenseByMethod().putAll(getExpenseByMethod(userId, year, month, dashboardDto.getTimeline()));

        dashboardDto.getFixatedExpenses().put("contas", getExpensesByCategory(new CategoryFilterDto(userId, year, month, "Conta")));
        dashboardDto.getFixatedExpenses().put("streaming", getExpensesByCategory(new CategoryFilterDto(userId, year, month, "Streaming")));

        return dashboardDto;
    }

    public List<ActivityDto> getExpensesByCategory(CategoryFilterDto categoryFilterDto) {

        List<Date> dates = dateUtils.calendarGenerator(categoryFilterDto.getYear(), categoryFilterDto.getMonth());

        Optional<List<Expense>> expenses = expenseRepository.findAllByUserIdAndPaidAtBetweenAndCategoryName(categoryFilterDto.getUserId(), dates.get(0), dates.get(1), categoryFilterDto.getName());

        if(expenses.isEmpty()) {
            return Collections.emptyList();
        }

        return activityExpenseMapper.toDto(expenses.get());
    }

    public HashMap<String, Double> getExpenseByCategory(String userId, String year, String month, List<ActivityDto> timeline) {
        HashMap<String, Double> result = new HashMap<>();

        for (ActivityDto activityDto : timeline) {
            if(activityDto.getType().equals("income")){
                continue;
            }
            else
            if(result.containsKey(activityDto.getActivityCategory())){
                result.put(activityDto.getActivityCategory(), Math.floor(result.get(activityDto.getActivityCategory()) + activityDto.getValue() * 100) / 100);
            }
            else{
                result.put(activityDto.getActivityCategory(), Math.floor(activityDto.getValue() * 100) / 100);
            }
        }
        return result;
    }

    public HashMap<String, Double> getExpenseByMethod(String userId, String year, String month,  List<ActivityDto> timeline) {
        HashMap<String, Double> result = new HashMap<>();

        for (ActivityDto activityDto : timeline) {
            if(activityDto.getType().equals("income")){
                continue;
            }
            else
            if(result.containsKey(activityDto.getMethod())){
                Double calc = Math.floor( (result.get(activityDto.getMethod()) + activityDto.getValue()) * 100) / 100;

                result.put(activityDto.getMethod(), calc);
            }
            else{
                result.put(activityDto.getMethod(), Math.floor(activityDto.getValue() * 100) / 100);
            }
        }
        return result;
    }

    public List<CategoryByMonthDto> getExpenseByCategoryByMonth(String userId, String year, String month) {
        List<CategoryByMonthDto> result = new ArrayList<>();

        return result;
    }

    public TotalsDto getTotals(String year, String month, String userId, String type){

        Dictionary<String, String> months = dateUtils.months();

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
                List<Date> calendar = dateUtils.calendarGenerator(year, month);

                Optional<List<Income>> value = incomeRepository.findAllByUserIdAndReceivedAtBetweenOrderByReceivedAt(userId, calendar.get(0), calendar.get(1));
                Double cont = value.get().stream().map(Income::getValue).reduce(0.0, Double::sum);

                totals.setValue(cont);

                totalsRepository.save(totals);

            }
            else if(Objects.equals(type, "expense")){
                List<Date> calendar = dateUtils.calendarGenerator(year, month);

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
            int monthAux = dateUtils.convertMonthToNumber(month) - i;
            String yearAux = year;

            if(monthAux < 1){
                monthAux = 12 + monthAux;
                yearAux = String.valueOf(Integer.parseInt(yearAux) - 1);
            }
            else if(monthAux > 12){
                monthAux = monthAux - 12;
                yearAux = String.valueOf(Integer.parseInt(yearAux) + 1);
            }

            totals.add(getTotals(yearAux, dateUtils.convertMonthToString(monthAux), userId, type));
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

    public List<IncomeVSExpense> getIncomeVsExpense(String userId, String year, String month){
        List<IncomeVSExpense> result = new ArrayList<>();

        int monthAux = dateUtils.convertMonthToNumber(month);

        String yearAux = year;

        for(int i = 0; i < 7; i++){
            if(monthAux < 1){
                monthAux = 12;
                yearAux = String.valueOf(Integer.parseInt(yearAux) - 1);
            }

            String monthString = dateUtils.convertMonthToString(monthAux);

            TotalsDto income = getTotals(yearAux, monthString, userId, "income");
            TotalsDto expense = getTotals(yearAux, monthString, userId, "expense");

            IncomeVSExpense novo = new IncomeVSExpense();

            novo.setMonthYear(monthString);
            novo.setIncome(income.getValue());
            novo.setExpense(expense.getValue());
            result.add(novo);

            monthAux--;
        }

        monthAux = dateUtils.convertMonthToNumber(month) + 1;
        yearAux = year;

        Collections.reverse(result);

        for (int i = 0; i < 6; i++) {
            if (monthAux > 12) {
                monthAux = 1;
                yearAux = String.valueOf(Integer.parseInt(yearAux) + 1);
            }
            String monthString = dateUtils.convertMonthToString(monthAux);

            TotalsDto income = getTotals(yearAux, monthString, userId, "income");
            TotalsDto expense = getTotals(yearAux, monthString, userId, "expense");

            IncomeVSExpense novo = new IncomeVSExpense();

            novo.setMonthYear(monthString);
            novo.setIncome(income.getValue());
            novo.setExpense(expense.getValue());
            result.add(novo);

            monthAux++;
        }

        return result;

    }

    public List<ActivityDto> timelineByMonth(String userId, String year, String month) {

        List<Date> calendar = dateUtils.calendarGenerator(year, month);

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
