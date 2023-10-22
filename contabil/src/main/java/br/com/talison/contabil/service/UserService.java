package br.com.talison.contabil.service;

import br.com.talison.contabil.model.User;
import br.com.talison.contabil.service.dto.UserDto;
import br.com.talison.contabil.repository.UserRepository;
import br.com.talison.contabil.service.mapper.UserMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Date;
import java.util.List;
import java.util.Optional;

@Service
@Transactional
@RequiredArgsConstructor
public class UserService {

    private final UserMapper userMapper = new UserMapper();
    private final UserRepository userRepository;


    public List<UserDto> listar() {
        return userMapper.toDTO(userRepository.findAll());
    }

    public String addUser(UserDto user) {
        user.setCreatedAt(new Date());
        return userRepository.save(userMapper.toEntity(user)).getId();
    }

    public UserDto updateUser(UserDto dto) {
        if (userRepository.existsById(dto.getId())) {
            userRepository.save(userMapper.toEntity(dto));
            Optional<User> novo = userRepository.findById(dto.getId());

            if(novo.isPresent()){
                return userMapper.toDTO(novo.get());
            }
        }
        return null;
    }

    public void deleteUser(String id) {
        userRepository.deleteById(id);
    }

    public UserDto getUserById(String id) {
        Optional<User> novo = userRepository.findById(id);
        return novo.map(userMapper::toDTO).orElse(null);
    }

}
