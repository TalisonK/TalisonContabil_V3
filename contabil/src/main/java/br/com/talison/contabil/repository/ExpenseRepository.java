package br.com.talison.contabil.repository;

import br.com.talison.contabil.domain.Expense;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.Optional;

public interface ExpenseRepository extends MongoRepository<Expense, String> {


    Optional<Expense> findByDescription(String description);

}
