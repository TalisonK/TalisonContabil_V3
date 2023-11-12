package br.com.talison.contabil.repository;

import br.com.talison.contabil.domain.Income;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.Date;
import java.util.List;
import java.util.Optional;

public interface IncomeRepository extends MongoRepository<Income, String> {

    Optional<List<Income>> findAllByUserIdAndReceivedAtBetweenOrderByReceivedAt(String user_id, Date receivedAt_start, Date receivedAt_end);

}
