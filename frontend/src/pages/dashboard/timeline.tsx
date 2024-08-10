import React from 'react'
import { DisplayFlex, Image, Text } from '../../styles'
import Activity from '../../interfaces/Activity'
import { Divider } from '@mui/material'
import {
    faArrowUpFromBracket,
    faArrowsDownToLine,
} from '@fortawesome/free-solid-svg-icons'

interface TimelineProps {
    activities: Activity[]
}

interface ItemProps {
    activity: Activity
}

const TimelineItem = ({ activity }: ItemProps) => {
    return (
        <DisplayFlex direction="row">
            {/* simbolos */}
            <DisplayFlex
                width="20px"
                height="100%"
                backgroundColor={
                    activity.type === 'income' ? '#048063' : '#db5151'
                }
            >
                {activity.type === 'income' ? (
                    <DisplayFlex width="20px">
                        <Image
                            src="income.png"
                            height="20px"
                            marginBottom="auto"
                            marginTop="2px"
                            marginLeft="1px"
                        />
                    </DisplayFlex>
                ) : (
                    <DisplayFlex width="20px">
                        <Image
                            src="expense.png"
                            height="20px"
                            marginBottom="auto"
                            marginTop="2px"
                            marginLeft="1px"
                        />
                    </DisplayFlex>
                )}
            </DisplayFlex>
            {/* linha */}
            <DisplayFlex direction="column" marginLeft="5px" width="100%">
                <Text fontSize="1.4em" margin="0">
                    {activity.description}
                </Text>
                {/* categoria */}
                <DisplayFlex direction="row" width="100%" height="100%">
                    <DisplayFlex width="25%" style={{ minWidth: '50px' }}>
                        <Text
                            fontSize="0.8em"
                            marginBottom="auto"
                            marginTop="2px"
                            marginLeft="5px"
                        >
                            {activity.categoryName}
                        </Text>
                    </DisplayFlex>
                    <DisplayFlex width="100%" justifyContent="space-evenly">
                        {/* valor */}
                        <DisplayFlex
                            direction="column"
                            justifyContent="center"
                            style={{
                                minWidth: '70px',
                            }}
                        >
                            <Text marginTop="3px" marginBottom="5px">
                                Valor
                            </Text>
                            <Text margin="0" marginBottom="10px">
                                R$ {activity.value}
                            </Text>
                        </DisplayFlex>

                        {/* parcelas */}
                        <DisplayFlex
                            direction="column"
                            justifyContent="center"
                            style={{ alignItems: 'center' }}
                        >
                            <Text
                                marginTop="3px"
                                marginBottom="5px"
                                style={{
                                    alignSelf: 'center',
                                    minWidth: '70px',
                                    alignItems: 'center',
                                    textAlign: 'center',
                                }}
                            >
                                {activity.paymentMethod === 'CREDIT_CARD'
                                    ? 'Parcelas'
                                    : activity.paymentMethod === 'DEBIT_CARD'
                                    ? 'Débito'
                                    : activity.paymentMethod === 'PIX'
                                    ? 'PIX'
                                    : ''}
                            </Text>
                            <Text
                                margin="0"
                                marginBottom="10px"
                                style={{ textAlign: 'center' }}
                            >
                                {activity.paymentMethod === 'CREDIT_CARD'
                                    ? `${activity.actualParcel} - ${activity.totalParcel}`
                                    : ''}
                            </Text>
                        </DisplayFlex>
                    </DisplayFlex>
                </DisplayFlex>
            </DisplayFlex>

            {/* data */}
            <DisplayFlex
                width="100px"
                marginLeft="auto"
                marginRight="5px"
                direction="column"
            >
                <Text fontSize="0.8em" margin="0">
                    {activity.activityDate.slice(0, 10)}
                </Text>
                <Text fontSize="0.8em" margin="0">
                    {activity.activityDate.slice(11)}
                </Text>
            </DisplayFlex>
        </DisplayFlex>
    )
}

const Timeline = (props: TimelineProps) => {
    return props.activities === null? 
    (<Text>Não há atividades</Text>) :
    (<DisplayFlex direction="column" overflow="auto" width="100%">
        {props.activities.map((activity) => (
            <>
                <TimelineItem activity={activity} />
                <Divider />
            </>
        ))}
    </DisplayFlex>)

}

export default Timeline
