import React from "react";
import { useSnackbar, VariantType } from "notistack";

const NotificationStructure = () => {

    const {enqueueSnackbar} = useSnackbar();

    const handleNotification = (message: string) => {
        enqueueSnackbar(message);
    }

    const handleNotificationVariant = (messagee: string, variant: VariantType) => {
        enqueueSnackbar(messagee, { variant });
    };
}

export default NotificationStructure;