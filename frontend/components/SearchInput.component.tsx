type SearchInputProps = {
  searchText: string;
  handleChange: (searchText: string) => void;
};

function SearchInputComponent({ searchText, handleChange }: SearchInputProps) {
  return (
    <input
      type="text"
      placeholder="search... cats"
      value={searchText}
      onChange={(e) => handleChange(e.target.value)}
    />
  );
}

export default SearchInputComponent;
