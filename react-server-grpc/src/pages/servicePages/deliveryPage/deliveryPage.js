import {
	Box,
	Button,
	FormControl,
	InputLabel,
	MenuItem,
	Select,
	Typography,
} from "@mui/material";
import React, {
	createContext,
	useCallback,
	useContext,
	useEffect,
} from "react";
import { useAuth } from "../../../hooks/auth.hook";
import { itemContext } from "./item.Context";
import DeliveryItemBlock from "./blocks/item.block";
import axios from "axios";
import { AuthContext } from "../../../context/auth.context";

const DeliveryPage = () => {
	const auth = useContext(AuthContext);
	const [deliveries, setDeliveries] = React.useState({
		provider: "",
	});
	const [providers, setProviders] = React.useState([]);
	const [items, setItems] = React.useState([]);
	const [selectedItems, setSelectedItems] = React.useState([]);
	const [itemsCount, setItemsCount] = React.useState(0);

	const handleChangeDeliver = (e) => {
		setDeliveries({
			...deliveries,
			[e.target.name]: e.target.value,
		});
	};
	const GetAllProviders = useCallback(async () => {
		await axios
			.get("http://localhost:8000/allproviders", {
				headers: { Authorization: `Bearer ${auth.token}` },
			})
			.then((res) => {
				setProviders(res.data.providers);
			});
	}, []);
	const GetAllItems = useCallback(async () => {
		await axios.get("http://localhost:8000/getallitems").then((res) => {
			setItems(res.data.allItems);
		});
	}, []);

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
						sx={{ width: "250px" }}
						label="Поставщик"
						id="provider"
						name="provider"
						onChange={handleChangeDeliver}
					>
						{providers.map((provider) => (
							<MenuItem value={provider.id}>
								{provider.name}
							</MenuItem>
						))}
					</Select>
				</FormControl>
				<Button>Добавить нового поставщика</Button>
			</Box>
			<Button>Добавить товар в накладную</Button>
			<Button>Удалить товар из накладной</Button>
			<itemContext.Provider value={[selectedItems, setSelectedItems]}>
				{selectedItems.map((item) => (
					<DeliveryItemBlock items={items} />
				))}
			</itemContext.Provider>
			<Box>
				<FormControl>
					<InputLabel>Товары</InputLabel>
					<Select
						sx={{ width: "250px" }}
						label="Товары"
						id="items"
						name="items"
						onChange={handleChangeDeliver}
					>
						{items.map((item) => (
							<MenuItem value={item.id}>{item.name}</MenuItem>
						))}
					</Select>
				</FormControl>
			</Box>
			<Box>
				<Button>Сформировать накладную</Button>
			</Box>
		</div>
	);
};

export default DeliveryPage;
