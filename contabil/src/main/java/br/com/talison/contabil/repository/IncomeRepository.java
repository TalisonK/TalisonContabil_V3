package br.com.talison.contabil.repository;

import br.com.talison.contabil.domain.Income;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface IncomeRepository extends MongoRepository<Income, String> {
}
