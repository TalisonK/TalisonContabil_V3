package br.com.talison.contabil.repository;

import br.com.talison.contabil.domain.Expense;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.Date;
import java.util.List;
import java.util.Optional;

public interface ExpenseRepository extends MongoRepository<Expense, String> {


    Optional<Expense> findByDescription(String description);

    Optional<List<Expense>> findAllByUserIdAndPaidAtBetweenOrderByPaidAt(String user_id, Date paidAt, Date paidAt2);

    Optional<List<Expense>> findAllByUserIdAndPaidAtBetweenAndCategoryName(String user_id, Date paidAt, Date paidAt2, String category_name);

    Optional<List<Expense>> findAllByUserId(String id);
}
