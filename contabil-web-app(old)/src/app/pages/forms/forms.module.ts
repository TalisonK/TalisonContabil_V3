import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {PrimengModule} from "../../shared/primeng.module";
import {CategoryServices} from "../../../services/Category.services";
import {MethodsUtil} from "../../shared/utils/methods.util";
import {ExpenseServices} from "../../../services/Expense.services";
import {IncomeServices} from "../../../services/Income.services";


@NgModule({
    declarations: [],
    imports: [
        CommonModule,
        PrimengModule

    ],
    providers: [
        CategoryServices,
        MethodsUtil,
        ExpenseServices,
        IncomeServices
    ]
})
export class FormsModule {
}
