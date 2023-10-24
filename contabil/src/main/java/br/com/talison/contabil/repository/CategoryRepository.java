package br.com.talison.contabil.repository;

import br.com.talison.contabil.domain.Category;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.Optional;

public interface CategoryRepository extends MongoRepository<Category, String> {

    Optional<Category> findByName(String categoryName);
}
