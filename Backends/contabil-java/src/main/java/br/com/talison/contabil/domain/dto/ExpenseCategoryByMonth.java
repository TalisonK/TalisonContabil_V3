package br.com.talison.contabil.domain.dto;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class ExpenseCategoryByMonth {

    private String category;

    private Double value;
}
