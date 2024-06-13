package br.com.talison.contabil.service.mapper;

import br.com.talison.contabil.domain.Totals;
import br.com.talison.contabil.domain.dto.TotalsDto;
import org.mapstruct.Mapper;
import org.mapstruct.ValueMapping;

@Mapper(componentModel = "spring")
public interface TotalsMapper extends EntityMapper<TotalsDto, Totals>{


}
