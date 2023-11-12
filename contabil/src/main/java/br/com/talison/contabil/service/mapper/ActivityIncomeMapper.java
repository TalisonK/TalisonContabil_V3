package br.com.talison.contabil.service.mapper;

import br.com.talison.contabil.domain.Income;
import br.com.talison.contabil.domain.dto.ActivityDto;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;

@Mapper(componentModel = "spring")
public interface ActivityIncomeMapper extends EntityMapper<ActivityDto, Income>{

    @Override
    @Mapping(source = "user.id", target = "userId")
    @Mapping(source = "receivedAt", target = "activityDate")
    ActivityDto toDto(Income entity);

    @Override
    @Mapping(source = "userId", target = "user.id")
    @Mapping(source = "activityDate", target = "receivedAt")
    Income toEntity(ActivityDto dto);

}
