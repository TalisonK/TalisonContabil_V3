package br.com.talison.contabil.service.mapper;

import br.com.talison.contabil.domain.Income;
import br.com.talison.contabil.domain.dto.IncomeDto;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;

@Mapper(componentModel = "spring")
public interface IncomeMapper extends EntityMapper<IncomeDto, Income>{

    @Override
    @Mapping(source = "user", target = "user.name")
    Income toEntity(IncomeDto dto);

    @Override
    @Mapping(source = "user.name", target = "user")
    IncomeDto toDto(Income entity);

}
