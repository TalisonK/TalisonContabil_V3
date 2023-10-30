package br.com.talison.contabil.service;

import br.com.talison.contabil.domain.Expense;
import br.com.talison.contabil.domain.Income;
import br.com.talison.contabil.domain.Totals;
import br.com.talison.contabil.domain.dto.TotalsDto;
import br.com.talison.contabil.repository.ExpenseRepository;
import br.com.talison.contabil.repository.IncomeRepository;
import br.com.talison.contabil.repository.TotalsRepository;
import br.com.talison.contabil.repository.UserRepository;
import br.com.talison.contabil.service.mapper.TotalsMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.*;

@Service
@Transactional
@RequiredArgsConstructor
public class TotalsService {

    private final TotalsRepository totalsRepository;
    private final IncomeRepository incomeRepository;
    private final ExpenseRepository expenseRepository;
    private final UserRepository userRepository;
    private final TotalsMapper totalsMapper;

    private List<Date> calendarGenerator(String year, String month) {
        String start = year + "-" + month + "-01";
        String end = year + "-" + month + "-31";

        SimpleDateFormat form = new SimpleDateFormat("yyyy-MM-dd");

        try{
            Date Date1 = form.parse(start);
            Date Date2 = form.parse(end);
            return Arrays.asList(Date1, Date2);
        }catch (ParseException e){
            e.printStackTrace();
        }
        return null;
    }
    private void createTotals(String year, String month, String userId){

        Optional<Totals> income = totalsRepository.findByTypeAndYearAndMonthAndUserId("income", year, month, userId);
        Optional<Totals> expense = totalsRepository.findByTypeAndYearAndMonthAndUserId("expense", year, month, userId);

        Totals totals = new Totals();

        Date date = null;

        try {
            date = new SimpleDateFormat("MMM", Locale.ENGLISH).parse(month);
        } catch (ParseException e) {
            e.printStackTrace();
            return;
        }

        Calendar cal = Calendar.getInstance();
        cal.setTime(date);
        cal.set(Calendar.YEAR, Integer.parseInt(year));


        if(income.isEmpty() || expense.isEmpty()){
            totals.setYear(year);
            totals.setMonth(month);
            totals.setUpdatedAt(new Date());
            totals.setCreatedAt(new Date());
            totals.setUserId(userId);
        }

        List<Date> calendar = calendarGenerator(year, String.valueOf(cal.get(Calendar.MONTH)+1));

        if(income.isEmpty()) {
            Optional<List<Income>> value = incomeRepository.findAllByUserIdAndReceivedAtBetween(userId, calendar.get(0), calendar.get(1));

            Double cont = value.get().stream().map(Income::getValue).reduce(0.0, Double::sum);

            totals.setType("income");
            totals.setValue(cont);
            totalsRepository.save(totals);
        }

        if(expense.isEmpty()){
            Optional<List<Expense>> value = expenseRepository.findAllByUserIdAndPaidAtBetween(userId, calendar.get(0), calendar.get(1));

            Double cont = value.get().stream().map(Expense::getValue).reduce(0.0, Double::sum);
            totals.setType("expense");
            totals.setValue(cont);
            totalsRepository.save(totals);
        }
    }

    public TotalsDto getTotals(String year, String month, String userId, String type){

        Optional<Totals> explorer = totalsRepository.findByTypeAndYearAndMonthAndUserId(type, year, month, userId);

        if(explorer.isEmpty()){
            createTotals(year, month, userId);
            return getTotals(year, month, userId, type);
        }
        else{
            return totalsMapper.toDto(explorer.get());
        }
    }

    public TotalsDto updateTotals(String year, String month, String userId, String type){

        Optional<Totals> explorer = totalsRepository.findByTypeAndYearAndMonthAndUserId(type, year, month, userId);

        if(explorer.isEmpty()){
            return null;
        }
        else{
            Totals totals = explorer.get();
            totals.setUpdatedAt(new Date());

            if(Objects.equals(type, "income")){
                List<Date> calendar = calendarGenerator(year, month);

                Optional<List<Income>> value = incomeRepository.findAllByUserIdAndReceivedAtBetween(userId, calendar.get(0), calendar.get(1));
                Double cont = value.get().stream().map(Income::getValue).reduce(0.0, Double::sum);

                totals.setValue(cont);

                totalsRepository.save(totals);

            }
            else if(Objects.equals(type, "expense")){
                List<Date> calendar = calendarGenerator(year, month);

                Optional<List<Expense>> value = expenseRepository.findAllByUserIdAndPaidAtBetween(userId, calendar.get(0), calendar.get(1));
                Double cont = value.get().stream().map(Expense::getValue).reduce(0.0, Double::sum);

                totals.setValue(cont);

                totalsRepository.save(totals);
            }

            return totalsMapper.toDto(totals);
        }
    }

}
