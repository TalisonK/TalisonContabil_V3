package br.com.talison.contabil.controller;


import br.com.talison.contabil.domain.dto.ExpenseDto;
import br.com.talison.contabil.service.ExpenseService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;
import java.util.List;

@RestController
@RequestMapping("/api/expense")
@RequiredArgsConstructor
public class ExpenseController {

    private final ExpenseService expenseService;

    @GetMapping("/all")
    public ResponseEntity<List<ExpenseDto>> getAllExpenses() {
        return ResponseEntity.ok(expenseService.listar());
    }

    @GetMapping("/{id}")
    public ResponseEntity<ExpenseDto> getExpenseById(@PathVariable String id) {
        ExpenseDto dto = expenseService.getExpenseById(id);

        if(dto != null) {
            return ResponseEntity.status(200).body(dto);
        }
        else {
            return ResponseEntity.status(204).build();
        }
    }

    @PostMapping
    public ResponseEntity<List<String>> addExpense(@Valid @RequestBody ExpenseDto expense) {
        List<String> results = expenseService.addExpense(expense);

        if(results != null){
            return ResponseEntity.status(201).body(results);
        }
        else{
            return ResponseEntity.status(409).body(null);
        }
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteExpense(@PathVariable String id) {
        expenseService.delete(id);
        return ResponseEntity.status(200).build();

    }

    @PutMapping
    public ResponseEntity<ExpenseDto> updateExpense(@PathVariable ExpenseDto dto) {
        ExpenseDto ret = expenseService.updateExpense(dto);

        if(ret != null) {
            return ResponseEntity.status(200).body(ret);
        }
        else {
            return ResponseEntity.status(204).build();
        }
    }

}
