package br.com.talison.contabil.repository;

import br.com.talison.contabil.domain.Income;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.Date;
import java.util.List;
import java.util.Optional;

public interface IncomeRepository extends MongoRepository<Income, String> {

    Optional<List<Income>> findAllByUserIdAndReceivedAtBetween(String user_id, Date receivedAt, Date receivedAt2);

    Optional<List<Income>> findAllByUserIdAndReceivedAt_YearAndReceivedAt_Month(String user_id, int receivedAt_year, int receivedAt_month);

}
