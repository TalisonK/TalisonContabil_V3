package br.com.talison.contabil.service;

import br.com.talison.contabil.model.User;
import br.com.talison.contabil.service.dto.UserDto;
import br.com.talison.contabil.repository.UserRepository;
import br.com.talison.contabil.service.mapper.UserMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@Transactional
@RequiredArgsConstructor
public class UserService {

    private final UserMapper userMapper;

    private final UserRepository userRepository;

    public void addUser(UserDto user) {

        //User novo = userMapper.toEntity(user);

        //System.out.println(userRepository.insert());
    }
//
//    public void updateUser() {
//    }
//
//    public void deleteUser() {
//    }
//
//    public List<User> getAllUsers() {
//        List<User> all = userRepository.findAll();
//        return all;
//    }
//
//    public void getUserById() {
//    }
}
