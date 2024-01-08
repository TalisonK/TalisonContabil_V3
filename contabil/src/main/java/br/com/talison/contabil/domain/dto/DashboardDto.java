package br.com.talison.contabil.domain.dto;

import lombok.Getter;
import lombok.Setter;

import java.io.Serializable;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

@Getter
@Setter
public class DashboardDto implements Serializable {

    private String userId;

    private String year;

    private String month;

    private LocalDateTime updatedAt;

    private List<ActivityDto> timeline;

    private HashMap<String, List<TotalsDto>> resumes;

    private List<IncomeVSExpense> incomeVSexpense;

    private HashMap<String, Double> ExpenseByCategory;

    private HashMap<String, Double> ExpenseByMethod;

    private List<CategoryByMonthDto> categoryByMonth;


    public DashboardDto(String userId, String year, String month) {

        this.userId = userId;
        this.year = year;
        this.month = month;
        this.updatedAt = LocalDateTime.now();
        this.timeline = new ArrayList<>();

        this.resumes = new HashMap<>();
        this.resumes.put("incomes", new ArrayList<>());
        this.resumes.put("expenses", new ArrayList<>());
        this.resumes.put("balances", new ArrayList<>());

        this.incomeVSexpense = new ArrayList<>();
        this.ExpenseByCategory = new HashMap<>();
        this.ExpenseByMethod = new HashMap<>();
        this.categoryByMonth = new ArrayList<>();

    }


}
