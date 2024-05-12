package br.com.talison.contabil.domain.dto;

import lombok.Getter;
import lombok.Setter;

import java.util.List;

@Getter
@Setter
public class ExpenseListDto extends ExpenseDto{

    private List<ListItemDto> list;

}
