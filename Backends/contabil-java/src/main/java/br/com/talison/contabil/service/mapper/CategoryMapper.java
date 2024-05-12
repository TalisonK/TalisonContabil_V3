package br.com.talison.contabil.service.mapper;

import br.com.talison.contabil.domain.Category;
import br.com.talison.contabil.domain.dto.CategoryDto;
import org.mapstruct.Mapper;

@Mapper(componentModel = "spring")
public interface CategoryMapper extends EntityMapper<CategoryDto, Category>{



}
