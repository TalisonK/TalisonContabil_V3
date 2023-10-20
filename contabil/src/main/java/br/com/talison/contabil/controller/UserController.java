package br.com.talison.contabil.controller;

import br.com.talison.contabil.service.dto.UserDto;
import br.com.talison.contabil.service.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;

@RestController
@RequestMapping("/user")
@RequiredArgsConstructor
public class UserController {

    private final UserService userService;

    @PostMapping
    public ResponseEntity<Void> addUser(@Valid @RequestBody UserDto userDto) {

        userService.addUser(userDto);

        return ResponseEntity.status(201).build();

    }

//    @GetMapping
//    public ResponseEntity<Void> updateUser() {
//        System.out.println("updateUser");
//        return ResponseEntity.status(200).build();
//    }

//    public void deleteUser() {
//        //TODO
//    }
//
//    @GetMapping("/all")
//    public ResponseEntity<List<User>> getAllUsers() {
//        return ResponseEntity.ok(userService.getAllUsers());
//    }
//
//    public void getUserById() {
//        //TODO
//    }

}
