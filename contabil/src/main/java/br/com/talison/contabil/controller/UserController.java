package br.com.talison.contabil.controller;

import br.com.talison.contabil.domain.dto.UserCreateDto;
import br.com.talison.contabil.domain.dto.UserDto;
import br.com.talison.contabil.service.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;
import java.util.List;

@RestController
@RequestMapping("/user")
@RequiredArgsConstructor
public class UserController {


    private final UserService userService;

    @PostMapping
    public ResponseEntity<String> addUser(@Valid @RequestBody UserCreateDto userDto) {

        UserDto dto = userService.addUser(userDto);

        if(dto != null){
            return ResponseEntity.status(201).body(dto.getId());
        }
        else{
            return ResponseEntity.status(409).body(null);
        }
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteUser(@PathVariable String id) {
        userService.deleteUser(id);
        return ResponseEntity.status(200).build();

    }

    @GetMapping("/all")
    public ResponseEntity<List<UserDto>> getAllUsers() {
        return ResponseEntity.ok(userService.listar());
    }

    @GetMapping("/{id}")
    public ResponseEntity<UserDto> getUserById(@PathVariable String id) {
        UserDto dto = userService.getUserById(id);

        if(dto != null) {
            return ResponseEntity.status(200).body(dto);
        }
        else {
            return ResponseEntity.status(204).build();
        }
    }

    @PutMapping
    public ResponseEntity<UserDto> updateUser(@Valid @RequestBody UserDto dto) {
        UserDto ret = userService.updateUser(dto);

        if(ret != null){
            return ResponseEntity.status(200).body(ret);
        }
        else{
            return ResponseEntity.status(404).body(null);
        }
    }

    @PostMapping("/login")
    public ResponseEntity<UserDto> login(@Valid @RequestBody UserDto dto) {
        UserDto ret = userService.login(dto);

        if(ret != null){
            return ResponseEntity.status(200).body(ret);
        }
        else{
            return ResponseEntity.status(404).body(null);
        }
    }

}
