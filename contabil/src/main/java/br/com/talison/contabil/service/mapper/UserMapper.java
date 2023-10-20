package br.com.talison.contabil.service.mapper;

import br.com.talison.contabil.model.User;
import br.com.talison.contabil.service.dto.UserDto;
import org.mapstruct.Mapper;
@Mapper(componentModel = "spring")
public interface UserMapper extends EntityMapper<UserDto, User>{

}
