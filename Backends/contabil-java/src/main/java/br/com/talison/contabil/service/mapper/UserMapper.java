package br.com.talison.contabil.service.mapper;

import br.com.talison.contabil.domain.User;
import br.com.talison.contabil.domain.dto.UserDto;
import org.mapstruct.Mapper;

@Mapper(componentModel = "spring")
public interface UserMapper extends EntityMapper<UserDto, User> {
}
