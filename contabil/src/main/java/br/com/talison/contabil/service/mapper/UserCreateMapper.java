package br.com.talison.contabil.service.mapper;

import br.com.talison.contabil.domain.User;
import br.com.talison.contabil.domain.dto.UserCreateDto;
import org.mapstruct.Mapper;

@Mapper(componentModel = "spring")
public interface UserCreateMapper extends EntityMapper<UserCreateDto, User>{

}
