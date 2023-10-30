package br.com.talison.contabil.service;

import br.com.talison.contabil.domain.Category;
import br.com.talison.contabil.domain.Expense;
import br.com.talison.contabil.domain.User;
import br.com.talison.contabil.domain.dto.ExpenseDto;
import br.com.talison.contabil.domain.dto.TotalsDto;
import br.com.talison.contabil.domain.enums.EnumPaymentMethod;
import br.com.talison.contabil.repository.CategoryRepository;
import br.com.talison.contabil.repository.ExpenseRepository;
import br.com.talison.contabil.repository.UserRepository;
import br.com.talison.contabil.service.mapper.ExpenseMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.*;

@Service
@Transactional
@RequiredArgsConstructor
public class ExpenseService {

    private final ExpenseRepository expenseRepository;
    private final CategoryRepository categoryRepository;
    private final UserRepository userRepository;
    private final ExpenseMapper expenseMapper;
    private final TotalsService totalsService;


    private Expense sendExpense(ExpenseDto expense, Category category, User user) {
        Expense novo = new Expense(
                expense.getDescription(),
                expense.getPaymentMethod(),
                category,
                expense.getValue(),
                user,
                expense.getPaidAt(),
                expense.getActualParcel(),
                expense.getTotalParcel());
        return expenseRepository.save(novo);
    }

    public List<ExpenseDto> listar() {
        return expenseMapper.toDto(expenseRepository.findAll());
    }

    public List<String> addExpense(ExpenseDto expense) {

        //validations
        if(expense.getPaymentMethod() == EnumPaymentMethod.CREDIT_CARD && (expense.getTotalParcel() == null || expense.getActualParcel() == null)) {
            return null;
        }

        if((expense.getTotalParcel() != null && expense.getActualParcel() != null) && (expense.getTotalParcel() < expense.getActualParcel())) {
            return null;
        }

        Optional<Category> category = categoryRepository.findByName(expense.getCategory());
        Optional<User> user = userRepository.findByName(expense.getUser());

        if(category.isEmpty() || user.isEmpty()) {
            return null;
        }

        if(expense.getPaymentMethod() == EnumPaymentMethod.CREDIT_CARD) {
            if(expense.getPaidAt().getDate() <= 10){
                expense.setPaidAt(new Date(expense.getPaidAt().getYear(), expense.getPaidAt().getMonth(), 15));
            } else {
                expense.setPaidAt(new java.sql.Date(expense.getPaidAt().getYear(), expense.getPaidAt().getMonth() + 1, 15));
            }

            List<String> results = new ArrayList<>();

            for(; expense.getActualParcel() <= expense.getTotalParcel(); expense.setActualParcel(expense.getActualParcel() + 1)) {
                results.add(sendExpense(expense, category.get(), user.get()).getId());
                TotalsDto d = totalsService.getTotals(expense.getPaidAt(), user.get().getId(), "expense");
                //totalsService.addToTotals(d, expense.getValue());
                expense.setPaidAt(new java.sql.Date(expense.getPaidAt().getYear(), expense.getPaidAt().getMonth() + 1, 15));
            }

            return results;

        } else {
            expense.setActualParcel(null);
            expense.setTotalParcel(null);
            List<String> results = List.of(sendExpense(expense, category.get(), user.get()).getId());

            totalsService.updateTotals(expense.getPaidAt(), user.get().getId(), "expense");
            return results;
        }
    }

    public ExpenseDto updateExpense(ExpenseDto dto) {
        if (expenseRepository.existsById(dto.getId())) {

            Optional<User> user = userRepository.findByName(dto.getUser());

            expenseRepository.save(expenseMapper.toEntity(dto));
            totalsService.updateTotals(dto.getPaidAt(), user.get().getId(), "expense");
            return dto;
        }
        return null;
    }

    public void delete(String id) {
        expenseRepository.deleteById(id);
    }


    public ExpenseDto getExpenseById(String id) {
        return expenseMapper.toDto(expenseRepository.findById(id).orElse(null));
    }
}
