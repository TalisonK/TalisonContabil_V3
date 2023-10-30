package br.com.talison.contabil.repository;

import br.com.talison.contabil.domain.Totals;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.Date;
import java.util.List;
import java.util.Optional;

public interface TotalsRepository extends MongoRepository<Totals, String> {

    Optional<Totals> findByTypeAndYearAndMonthAndUserId(String type, String year, String month, String user);


}
