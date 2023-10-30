package br.com.talison.contabil.service;

import br.com.talison.contabil.domain.Income;
import br.com.talison.contabil.domain.User;
import br.com.talison.contabil.domain.dto.IncomeDto;
import br.com.talison.contabil.repository.IncomeRepository;
import br.com.talison.contabil.repository.UserRepository;
import br.com.talison.contabil.service.mapper.IncomeMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.Optional;

@Service
@Transactional
@RequiredArgsConstructor
public class IncomeService {

    private final IncomeRepository incomeRepository;
    private final UserRepository userRepository;
    private final IncomeMapper incomeMapper;
    private final TotalsService totalsService;

    public List<IncomeDto> listar() {
        return incomeMapper.toDto(incomeRepository.findAll());
    }

    public IncomeDto addIncome(IncomeDto income) {

        Optional<User> user = userRepository.findByName(income.getUser());

        if(user.isEmpty()) {
            return null;
        }

        Income novo = new Income(
                income.getDescription(),
                income.getValue(),
                user.get(),
                income.getReceivedAt());

        totalsService.updateTotals(income.getReceivedAt(), user.get().getId(), "expense");

        return incomeMapper.toDto(incomeRepository.save(novo));
    }

    public IncomeDto updateIncome(IncomeDto dto) {
        if (incomeRepository.existsById(dto.getId())) {
            incomeRepository.save(incomeMapper.toEntity(dto));

            User user = userRepository.findByName(dto.getUser()).get();

            totalsService.updateTotals(dto.getReceivedAt(), user.getId(), "expense");
            return dto;
        }
        return null;
    }

    public void delete(String id) {
        incomeRepository.deleteById(id);
    }

    public IncomeDto getIncomeById(String id) {
        return incomeMapper.toDto(incomeRepository.findById(id).orElse(null));
    }
}
