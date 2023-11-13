package br.com.talison.contabil.controller;

import br.com.talison.contabil.domain.dto.CategoryDto;
import br.com.talison.contabil.service.CategoryService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;
import java.util.List;

@RestController
@RequestMapping("/category")
@RequiredArgsConstructor
public class CategoryController {

    private final CategoryService categoryService;

    @GetMapping("/all")
    public ResponseEntity<List<CategoryDto>> getAllCategories() {
        return ResponseEntity.ok(categoryService.list());
    }

    @GetMapping("/{id}")
    public ResponseEntity<CategoryDto> getCategoryById(@PathVariable String id) {
        CategoryDto dto = categoryService.getCategoryById(id);

        if(dto != null) {
            return ResponseEntity.status(200).body(dto);
        }
        else {
            return ResponseEntity.status(204).build();
        }
    }

    @PostMapping
    public ResponseEntity<String> addCategory(@Valid @RequestBody CategoryDto category) {
        CategoryDto dto = categoryService.addCategory(category);

        if(dto != null){
            return ResponseEntity.status(201).body(dto.getId());
        }
        else{
            return ResponseEntity.status(409).body(null);
        }
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteCategory(@PathVariable String id) {
        categoryService.delete(id);
        return ResponseEntity.status(200).build();

    }

    @PutMapping
    public ResponseEntity<CategoryDto> updateCategory(@Valid @RequestBody CategoryDto dto) {
        CategoryDto ret = categoryService.updateCategory(dto);

        if(ret != null) {
            return ResponseEntity.status(200).body(ret);
        }
        else {
            return ResponseEntity.status(204).build();
        }
    }
}
