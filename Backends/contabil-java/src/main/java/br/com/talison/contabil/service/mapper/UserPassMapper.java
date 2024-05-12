package br.com.talison.contabil.service.mapper;

import br.com.talison.contabil.domain.User;
import br.com.talison.contabil.domain.dto.UserPassDto;
import org.mapstruct.Mapper;

@Mapper(componentModel = "spring")
public interface UserPassMapper extends EntityMapper<UserPassDto, User>{

}
