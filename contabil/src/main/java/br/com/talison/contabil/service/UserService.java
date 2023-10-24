package br.com.talison.contabil.service;

import br.com.talison.contabil.domain.User;
import br.com.talison.contabil.domain.dto.UserCreateDto;
import br.com.talison.contabil.domain.dto.UserDto;
import br.com.talison.contabil.repository.UserRepository;
import br.com.talison.contabil.service.mapper.UserCreateMapper;
import br.com.talison.contabil.service.mapper.UserMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Example;
import org.springframework.data.domain.ExampleMatcher;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Date;
import java.util.List;
import java.util.Optional;

@Service
@Transactional
@RequiredArgsConstructor
public class UserService {

    private final UserMapper userMapper;
    private final UserCreateMapper userCreateMapper;
    private final UserRepository userRepository;


    public List<UserDto> listar() {
        return userMapper.toDto(userRepository.findAll());
    }

    public UserDto addUser(UserCreateDto user) {

        User novo = new User();
        novo.setName(user.getName());

        Example<User> example = Example.of(novo, ExampleMatcher.matching().withIgnorePaths("id", "createdAt", "role", "password"));

        Optional<User> verification = userRepository.findOne(example);

        if(verification.isPresent()) {
            return null;
        }

        user.setCreatedAt(new Date());
        User entrada = userRepository.save(userCreateMapper.toEntity(user));

        return userMapper.toDto(entrada);
    }

    public UserDto updateUser(UserDto dto) {
        if (userRepository.existsById(dto.getId())) {
            userRepository.save(userMapper.toEntity(dto));
            Optional<User> novo = userRepository.findById(dto.getId());

            if(novo.isPresent()){
                return userMapper.toDto(novo.get());
            }
        }
        return null;
    }

    public void deleteUser(String id) {
        userRepository.deleteById(id);
    }

    public UserDto getUserById(String id) {
        Optional<User> novo = userRepository.findById(id);
        return novo.map(userMapper::toDto).orElse(null);
    }

    public UserDto login(UserDto dto) {
        User user = userMapper.toEntity(dto);
        Example<User> example = Example.of(user, ExampleMatcher.matching().withIgnorePaths("id", "createdAt", "role", "name"));

        Optional<User> verification = userRepository.findOne(example);

        return verification.map(userMapper::toDto).orElse(null);
    }
}
