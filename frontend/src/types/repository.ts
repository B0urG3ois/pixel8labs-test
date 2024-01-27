export interface Repository {
  id: number;
  name: string;
  full_name: string;
  private: boolean;
  description: string | null;
  topics: string[];
  language: string;
  url: string;
  updated_at: string;
}