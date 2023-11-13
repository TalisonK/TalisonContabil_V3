package br.com.talison.contabil.service;

import br.com.talison.contabil.domain.dto.CategoryDto;
import br.com.talison.contabil.repository.CategoryRepository;
import br.com.talison.contabil.service.mapper.CategoryMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
@Transactional
@RequiredArgsConstructor
public class CategoryService {

    private final CategoryRepository categoryRepository;
    private final CategoryMapper categoryMapper;

    public List<CategoryDto> list() {
        return categoryMapper.toDto(categoryRepository.findAll());
    }

    public CategoryDto addCategory(CategoryDto category) {
        return categoryMapper.toDto(categoryRepository.save(categoryMapper.toEntity(category)));
    }

    public CategoryDto updateCategory(CategoryDto dto) {
        if (categoryRepository.existsById(dto.getId())) {
            categoryRepository.save(categoryMapper.toEntity(dto));
            return dto;
        }
        return null;
    }

    public void delete(String id) {
        categoryRepository.deleteById(id);
    }

    public CategoryDto getCategoryById(String id) {
        return categoryMapper.toDto(categoryRepository.findById(id).orElse(null));
    }
}
