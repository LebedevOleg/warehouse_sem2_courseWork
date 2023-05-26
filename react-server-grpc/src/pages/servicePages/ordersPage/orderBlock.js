import { Button, Stack, Typography } from "@mui/material";
import React from "react";

const OrderBlock = ({ order }) => {
	return (
		<Stack spacing={2} direction="row">
			<Typography>{order.id}</Typography>
			<Typography>{order.date_start}</Typography>
			{(order.date_end.Valid && "Не закрыт") || (
				<Typography>{order.date_end.String}</Typography>
			)}

			<Typography>{order.status}</Typography>
			<Typography>{order.price}</Typography>
			<Typography>{order.address}</Typography>
			<Typography>{order.user_id}</Typography>
			<Button>Подробнее</Button>
			<Button>Удалить</Button>
		</Stack>
	);
};

export default OrderBlock;
