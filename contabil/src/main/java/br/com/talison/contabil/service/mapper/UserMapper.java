package br.com.talison.contabil.service.mapper;

import br.com.talison.contabil.model.EnumRole;
import br.com.talison.contabil.model.User;
import br.com.talison.contabil.service.dto.UserDto;
import org.mapstruct.Mapper;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class UserMapper {

    public User toEntity (UserDto dto) {
        User novo = new User();
        novo.setId(dto.getId());
        novo.setName(dto.getName());
        novo.setPassword(dto.getPassword());
        novo.setRole(EnumRole.valueOf(dto.getRole()));
        novo.setCreatedAt(dto.getCreatedAt());

        return novo;
    }

    public List<User> toEntity(List<UserDto> lista) {
        return lista.stream().map(this::toEntity).collect(Collectors.toList());
    }

    public UserDto toDTO (User user) {
        UserDto novo = new UserDto();
        novo.setId(user.getId());
        novo.setName(user.getName());
        novo.setPassword(user.getPassword());
        novo.setRole(user.getRole().toString());
        novo.setCreatedAt(user.getCreatedAt());

        return novo;
    }

    public List<UserDto> toDTO(List<User> lista) {
        return lista.stream().map(this :: toDTO).collect(Collectors.toList());
    }

}
