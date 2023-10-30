package br.com.talison.contabil.repository;

import br.com.talison.contabil.domain.Expense;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.Date;
import java.util.List;
import java.util.Optional;

public interface ExpenseRepository extends MongoRepository<Expense, String> {


    Optional<Expense> findByDescription(String description);

    Optional<List<Expense>> findAllByUserIdAndPaidAtBetween(String user_id, Date paidAt, Date paidAt2);

    Optional<List<Expense>> findAllByUserIdAndPaidAt_YearAndPaidAt_Month(String user_id, int paidAt_year, int paidAt_month);

}
