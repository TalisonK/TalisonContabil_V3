package br.com.talison.contabil.repository;

import br.com.talison.contabil.domain.List;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface ListRepository extends MongoRepository<List, String> {
}
