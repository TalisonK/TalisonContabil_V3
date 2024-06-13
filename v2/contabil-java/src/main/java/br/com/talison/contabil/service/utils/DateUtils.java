package br.com.talison.contabil.service.utils;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.*;

public class DateUtils {

    public List<String> months = Arrays.asList("Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep","Oct","Nov","Dec");

    public Dictionary<String, String> months(){
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
    public Integer convertMonthToNumber(String month) {

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

    public String convertMonthToString(Integer month) {
        return months.get(month - 1);
    }


    public List<Date> calendarGenerator(String year, String month) {

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
}
