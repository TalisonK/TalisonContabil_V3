package br.com.talison.contabil.service;

import br.com.talison.contabil.domain.Income;
import br.com.talison.contabil.domain.User;
import br.com.talison.contabil.domain.dto.ActivityDto;
import br.com.talison.contabil.domain.dto.IncomeDto;
import br.com.talison.contabil.repository.IncomeRepository;
import br.com.talison.contabil.repository.UserRepository;
import br.com.talison.contabil.service.mapper.ActivityIncomeMapper;
import br.com.talison.contabil.service.mapper.IncomeMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Collections;
import java.util.Date;
import java.util.List;
import java.util.Optional;

@Service
@Transactional
@RequiredArgsConstructor
public class IncomeService {

    private final IncomeRepository incomeRepository;
    private final UserRepository userRepository;
    private final IncomeMapper incomeMapper;
    private final ActivityIncomeMapper activityIncomeMapper;
    private final TotalsService totalsService;

    public List<IncomeDto> list() {
        return incomeMapper.toDto(incomeRepository.findAll());
    }

    public List<ActivityDto> listActivities(String id) {

        Optional<List<Income>> incomes = incomeRepository.findAllByUserId(id);

        if(incomes.isEmpty()) {
            return Collections.emptyList();
        }

        List<ActivityDto> data = activityIncomeMapper.toDto(incomes.get());

        data = data.stream().map((dto) -> {
            dto.setType("Income");
            return dto;
        }).toList();

        return data;
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

        Income response = incomeRepository.save(novo);

        totalsService.updateTotals(income.getReceivedAt(), user.get().getId(), "income");

        return incomeMapper.toDto(response);
    }

    public IncomeDto updateIncome(IncomeDto dto) {
        if (incomeRepository.existsById(dto.getId())) {
            incomeRepository.save(incomeMapper.toEntity(dto));

            User user = userRepository.findByName(dto.getUser()).get();

            totalsService.updateTotals(dto.getReceivedAt(), user.getId(), "income");
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

