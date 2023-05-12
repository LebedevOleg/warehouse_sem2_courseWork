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
import FullFeaturedCrudGrid from "./blocks/item.block";

const DeliveryPage = () => {
	const auth = useContext(AuthContext);
	const [deliveries, setDeliveries] = React.useState({
		provider: 0,
		storage: 0,
	});
	const [providers, setProviders] = React.useState([]);
	const [storages, setStorages] = React.useState([]);

	const [items, setItems] = React.useState([]);
	const [selectedItems, setSelectedItems] = React.useState([]);

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
	const GetAllStorages = useCallback(async () => {
		await axios
			.get("http://localhost:8000/allstorages", {
				headers: { Authorization: `Bearer ${auth.token}` },
			})
			.then((res) => {
				setStorages(res.data.storages);
			});
	}, []);

	//todo: Бэк для создания накладной
	const handleCreateDelivery = async () => {
		await axios.post("http://localhost:8000/adddelivery", {
			provider: deliveries.provider,
			storage: deliveries.storage,
			items: selectedItems,
		});
	};

	useEffect(() => {
		GetAllProviders();
		GetAllItems();
		GetAllStorages();
	}, [GetAllProviders, GetAllItems, GetAllStorages]);

	return (
		<div>
			<Typography variant="h3">Сформировать накладную</Typography>
			<Box>
				<FormControl sx={{ mr: 1 }}>
					<InputLabel>Поставщик</InputLabel>
					<Select
						sx={{ width: "250px" }}
						label="Поставщик"
						id="provider"
						name="provider"
						onChange={handleChangeDeliver}
					>
						{providers.map((provider) => (
							<MenuItem value={provider.id} key={provider.id}>
								{provider.name}
							</MenuItem>
						))}
					</Select>
				</FormControl>
				<FormControl>
					<InputLabel>Склад</InputLabel>
					<Select
						sx={{ width: "250px" }}
						label="Склад"
						id="storage"
						name="storage"
						onChange={handleChangeDeliver}
					>
						{storages.map((storage) => (
							<MenuItem value={storage.id} key={storage.id}>
								{storage.name}
							</MenuItem>
						))}
					</Select>
				</FormControl>
			</Box>
			<Button>Добавить нового поставщика</Button>
			<itemContext.Provider value={[selectedItems, setSelectedItems]}>
				<FullFeaturedCrudGrid items={items} />
			</itemContext.Provider>
			<Box>
				<Button onClick={handleCreateDelivery}>
					Сформировать накладную
				</Button>
			</Box>
		</div>
	);
};

export default DeliveryPage;
