package br.com.talison.contabil.domain.dto;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
public class CategoryFilterDto {

    private String userId;

    private String name;

    private String month;

    private String year;


    public CategoryFilterDto(String userId, String year, String month, String conta) {
        this.userId = userId;
        this.year = year;
        this.month = month;
        this.name = conta;
    }
}
