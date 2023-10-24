package br.com.talison.contabil.service.mapper;

import br.com.talison.contabil.domain.Expense;
import br.com.talison.contabil.domain.dto.ExpenseDto;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;

@Mapper(componentModel = "spring")
public interface ExpenseMapper extends EntityMapper<ExpenseDto, Expense> {

    @Override
    @Mapping(source = "user", target = "user.name")
    @Mapping(source = "category", target = "category.name")
    Expense toEntity(ExpenseDto dto);

    @Override
    @Mapping(source = "user.name", target = "user")
    @Mapping(source = "category.name", target = "category")
    ExpenseDto toDto(Expense entity);

}
