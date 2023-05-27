import { Button, Stack, Typography } from "@mui/material";
import axios from "axios";
import React, { useContext } from "react";
import { AuthContext } from "../../../context/auth.context";

const OrderBlock = ({ order }) => {
	const auth = useContext(AuthContext);
	const [winOrder, setWinOrder] = React.useState(order);

	const handleUpdateStatus = async () => {
		await axios.post("http://localhost:8000/updateorderstatus", winOrder, {
			headers: { Authorization: `Bearer ${auth.token}` },
		});
	};
	return (
		<Stack spacing={2} direction="row">
			<Typography>{winOrder.id}</Typography>
			<Typography>{winOrder.date_start}</Typography>
			{(winOrder.date_end.Valid && "Не закрыт") || (
				<Typography>{winOrder.date_end.String}</Typography>
			)}

			<Typography>{winOrder.status}</Typography>
			<Typography>{winOrder.price}</Typography>
			<Typography>{winOrder.address}</Typography>
			<Typography>{winOrder.user_id}</Typography>
			<Button>Подробнее</Button>
			<Button onClick={handleUpdateStatus}>Закрыть</Button>
			<Button>Удалить</Button>
		</Stack>
	);
};

export default OrderBlock;
