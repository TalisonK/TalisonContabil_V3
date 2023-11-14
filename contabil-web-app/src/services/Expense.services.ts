import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";


@Injectable()

export class ExpenseServices{

    url = 'http://localhost:8080/expense';

    constructor(private http: HttpClient) {
    }


    createExpense(expense: any) {
        return this.http.post(this.url, expense);
    }

}
