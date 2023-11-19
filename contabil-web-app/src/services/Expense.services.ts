import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {environment} from "../environments/environment";


@Injectable()

export class ExpenseServices{

    url = environment.API_BASE_URL + '/expense';

    constructor(private http: HttpClient) {
    }


    createExpense(expense: any) {
        return this.http.post(this.url, expense);
    }

}
