/**
 * Generated by orval v7.2.0 🍺
 * Do not edit manually.
 * todo-schema-driven
 * OpenAPI spec version: 1.0
 */
export type GetTodoParams = {
todo_id?: number;
offset?: number;
limit?: number;
};

export interface TodoOpeError {
  /** エラーコード */
  error_code?: number;
  /** エラーメッセージ */
  error_message?: string;
}

export interface TodoId {
  /** TODOの登録番号 */
  todo_id?: number;
}

export interface TodoInformation {
  /**
   * TODOの詳細
   * @maxLength 300
   */
  detail?: string;
  /** 終了予定日 */
  end_date?: string;
  /**
   * TODOの進捗率
   * @minimum 0
   * @maximum 100
   */
  progress: number;
  /** 開始予定日 */
  start_date?: string;
  /**
   * TODOの件名
   * @minLength 1
   * @maxLength 100
   */
  title: string;
}

