import {
	Box,
	Button,
	FormControl,
	InputLabel,
	Select,
	Typography,
} from "@mui/material";
import React, { useCallback, useEffect } from "react";
import { useAuth } from "../../../hooks/auth.hook";

const DeliveryPage = () => {
	const { token } = useAuth();
	const [deliveries, setDeliveries] = React.useState({
		provider: "",
	});
	const [providers, setProviders] = React.useState([]);
	const [items, setItems] = React.useState([]);

	const handleChangeDeliver = (e) => {
		setDeliveries({
			...deliveries,
			[e.target.name]: e.target.value,
		});
	};
	const GetAllProviders = useCallback(async () => {}, []);
	const GetAllItems = useCallback(async () => {}, []);

	useEffect(() => {
		GetAllProviders();
		GetAllItems();
	}, [GetAllProviders, GetAllItems]);

	return (
		<div>
			<Typography variant="h3">Сформировать накладную</Typography>
			<Box>
				<FormControl>
					<InputLabel>Поставщик</InputLabel>
					<Select
						label="Поставщик"
						id="provider"
						name="provider"
						onChange={handleChangeDeliver}
					></Select>
				</FormControl>
				<Button>Добавить нового поставщика</Button>
			</Box>
			<Box>
				<FormControl>
					<InputLabel>Товары</InputLabel>
					<Select
						label="Товары"
						id="items"
						name="items"
						onChange={handleChangeDeliver}
					></Select>
				</FormControl>
			</Box>
			<Box>
				<Button>Сформировать накладную</Button>
			</Box>
		</div>
	);
};

export default DeliveryPage;
