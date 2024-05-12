package br.com.talison.contabil.domain.dto;


import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
public class IncomeVSExpense {
    private String monthYear; //"Dec-2023",
    private Double income; //3500

    private Double expense; //4000

}
