package br.com.talison.contabil.controller;

import br.com.talison.contabil.service.dto.UserDto;
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
    public ResponseEntity<String> addUser(@Valid @RequestBody UserDto userDto) {

        return ResponseEntity.status(201).body(userService.addUser(userDto));

    }

    public void deleteUser() {

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

}
