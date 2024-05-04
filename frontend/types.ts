
export interface SearchData {
  id: string;
  url: string;
  title: string;
  related_image_tags: string[];
  annotated_image_tags: string[];
}

export interface SearchResponse {
  Query: string;
  CorrectedQuery: string;
  Documents: SearchData[];
}