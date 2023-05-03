import * as React from "react";
import PropTypes from "prop-types";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import AddIcon from "@mui/icons-material/Add";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/DeleteOutlined";
import SaveIcon from "@mui/icons-material/Save";
import CancelIcon from "@mui/icons-material/Close";
import {
	GridRowModes,
	DataGrid,
	GridToolbarContainer,
	GridActionsCellItem,
	useGridApiContext,
} from "@mui/x-data-grid";
import { itemContext } from "../item.Context";
import { Select } from "@mui/material";

function SelectEditInputCell(props) {
	const { id, value, field } = props;
	const apiRef = useGridApiContext();

	const handleChange = async (event) => {
		await apiRef.current.setEditCellValue({
			id,
			field,
			value: event.target.value,
		});
		apiRef.current.stopCellEditMode({ id, field });
	};

	return (
		<Select
			value={value}
			onChange={handleChange}
			size="small"
			sx={{ height: 1 }}
			native
			autoFocus
		>
			<option>Back-end Developer</option>
			<option>Front-end Developer</option>
			<option>UX Designer</option>
			{/* {items.map((item) => (
				<option key={item.id} value={item.name}>
					{" "}
					{item.name}{" "}
				</option>
			))} */}
		</Select>
	);
}

SelectEditInputCell.propTypes = {
	/**
	 * The column field of the cell that triggered the event.
	 */
	field: PropTypes.string.isRequired,
	/**
	 * The grid row id.
	 */
	id: PropTypes.oneOfType([PropTypes.number, PropTypes.string]).isRequired,
	/**
	 * The cell value.
	 * If the column has `valueGetter`, use `params.row` to directly access the fields.
	 */
	value: PropTypes.any,
};

const renderSelectEditInputCell = (params) => {
	return <SelectEditInputCell {...params} />;
};

function EditToolbar(props) {
	const { setRows, setRowModesModel, rows } = props;
	const handleClick = () => {
		const id = rows.length + 1;
		setRows((oldRows) => [...oldRows, { id: id, name: "", count: 0 }]);
		setRowModesModel((oldModel) => ({
			...oldModel,
			[id]: { mode: GridRowModes.Edit, fieldToFocus: "name" },
		}));
	};

	return (
		<GridToolbarContainer>
			<Button
				color="primary"
				startIcon={<AddIcon />}
				onClick={handleClick}
			>
				Добавить предмет
			</Button>
		</GridToolbarContainer>
	);
}

EditToolbar.propTypes = {
	setRowModesModel: PropTypes.func.isRequired,
	setRows: PropTypes.func.isRequired,
};

export default function FullFeaturedCrudGrid(props) {
	const [rows, setRows] = React.useContext(itemContext);
	const [rowModesModel, setRowModesModel] = React.useState({});

	const handleRowEditStart = (params, event) => {
		event.defaultMuiPrevented = true;
	};

	const handleRowEditStop = (params, event) => {
		event.defaultMuiPrevented = true;
	};

	const handleEditClick = (id) => () => {
		setRowModesModel({
			...rowModesModel,
			[id]: { mode: GridRowModes.Edit },
		});
	};

	const handleSaveClick = (id, name, count) => () => {
		setRowModesModel({
			...rowModesModel,
			[id]: { mode: GridRowModes.View },
		});
		console.log(id, name, count);
	};

	const handleDeleteClick = (id) => () => {
		setRows(rows.filter((row) => row.id !== id));
	};

	const handleCancelClick = (id) => () => {
		setRowModesModel({
			...rowModesModel,
			[id]: { mode: GridRowModes.View, ignoreModifications: true },
		});

		const editedRow = rows.find((row) => row.id === id);
		if (editedRow.isNew) {
			setRows(rows.filter((row) => row.id !== id));
		}
	};

	const processRowUpdate = (newRow) => {
		const updatedRow = { ...newRow, isNew: false };
		setRows(rows.map((row) => (row.id === newRow.id ? updatedRow : row)));
		return updatedRow;
	};

	const handleRowModesModelChange = (newRowModesModel) => {
		setRowModesModel(newRowModesModel);
	};

	const columns = [
		{
			field: "name",
			headerName: "Название товара",
			renderEditCell: renderSelectEditInputCell,
			width: 180,
			editable: true,
		},
		{
			field: "count",
			headerName: "Колличество",
			type: "number",
			editable: true,
		},
		{
			field: "actions",
			type: "actions",
			headerName: "Действия",
			width: 100,
			cellClassName: "actions",
			getActions: ({ id }) => {
				const isInEditMode =
					rowModesModel[id]?.mode === GridRowModes.Edit;

				if (isInEditMode) {
					return [
						<GridActionsCellItem
							icon={<SaveIcon />}
							label="Save"
							onClick={handleSaveClick(id)}
						/>,
						<GridActionsCellItem
							icon={<CancelIcon />}
							label="Cancel"
							className="textPrimary"
							onClick={handleCancelClick(id)}
							color="inherit"
						/>,
					];
				}

				return [
					<GridActionsCellItem
						icon={<EditIcon />}
						label="Edit"
						className="textPrimary"
						onClick={handleEditClick(id)}
						color="inherit"
					/>,
					<GridActionsCellItem
						icon={<DeleteIcon />}
						label="Delete"
						onClick={handleDeleteClick(id)}
						color="inherit"
					/>,
				];
			},
		},
	];

	return (
		<Box
			sx={{
				height: 500,
				width: "100%",
				"& .actions": {
					color: "text.secondary",
				},
				"& .textPrimary": {
					color: "text.primary",
				},
			}}
		>
			<DataGrid
				rows={rows}
				items={props.items}
				columns={columns}
				editMode="row"
				rowModesModel={rowModesModel}
				onRowModesModelChange={handleRowModesModelChange}
				onRowEditStart={handleRowEditStart}
				onRowEditStop={handleRowEditStop}
				processRowUpdate={processRowUpdate}
				slots={{
					toolbar: EditToolbar,
				}}
				slotProps={{
					toolbar: { setRows, setRowModesModel, rows },
				}}
			/>
		</Box>
	);
}
