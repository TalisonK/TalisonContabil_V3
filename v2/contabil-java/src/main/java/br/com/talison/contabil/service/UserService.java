package br.com.talison.contabil.service;

import br.com.talison.contabil.domain.User;
import br.com.talison.contabil.domain.dto.UserPassDto;
import br.com.talison.contabil.domain.dto.UserDto;
import br.com.talison.contabil.domain.enums.EnumRole;
import br.com.talison.contabil.repository.UserRepository;
import br.com.talison.contabil.service.mapper.UserPassMapper;
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
    private final UserPassMapper userPassMapper;
    private final UserRepository userRepository;


    public List<UserDto> listar() {
        return userMapper.toDto(userRepository.findAll());
    }

    public Boolean existsByName(String name) {
        return userRepository.existsByName(name);
    }

    public UserDto addUser(UserPassDto user) {

        if(existsByName(user.getName())) {
            return null;
        }
        User novo = new User();
        novo.setName(user.getName());

        Example<User> example = Example.of(novo, ExampleMatcher.matching().withIgnorePaths("id", "createdAt", "role", "password"));

        Optional<User> verification = userRepository.findOne(example);

        if(verification.isPresent()) {
            return null;
        }

        user.setCreatedAt(new Date());
        user.setRole(EnumRole.ROLE_USER.getValue());
        User entrada = userRepository.save(userPassMapper.toEntity(user));

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

    public UserDto login(UserPassDto dto) {
        User user = userPassMapper.toEntity(dto);

        Optional<User> novo = userRepository.findByNameAndPassword(user.getName(), user.getPassword());

        return novo.map(userMapper::toDto).orElse(null);

    }
}
