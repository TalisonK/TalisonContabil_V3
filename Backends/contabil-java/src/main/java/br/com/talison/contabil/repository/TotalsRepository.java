package br.com.talison.contabil.repository;

import br.com.talison.contabil.domain.Totals;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.Date;
import java.util.List;
import java.util.Optional;

public interface TotalsRepository extends MongoRepository<Totals, String> {

    Optional<Totals> findByTypeAndYearAndMonthAndUserId(String type, String year, String month, String user);

    Optional<Totals> findByUserIdAndYearAndMonthAndType(String user, String year, String month, String type);

    Optional<List<Totals>> findAllByTypeAndUserIdOrderByYearDesc(String type, String user);

    Optional<List<Totals>> findAllByUserIdAndYear(String user, String year);


}
