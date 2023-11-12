package br.com.talison.contabil.service.mapper;

import br.com.talison.contabil.domain.Expense;
import br.com.talison.contabil.domain.dto.ActivityDto;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;

@Mapper(componentModel = "spring")
public interface ActivityExpenseMapper extends EntityMapper<ActivityDto, Expense>{

    @Override
    @Mapping(source = "user.id", target = "userId")
    @Mapping(source = "paidAt", target = "activityDate")
    @Mapping(source="paymentMethod.", target="method")
    ActivityDto toDto(Expense entity);

    @Override
    @Mapping(source = "userId", target = "user.id")
    @Mapping(source = "activityDate", target = "paidAt")
    @Mapping(source="method", target="paymentMethod.")
    Expense toEntity(ActivityDto dto);

}
