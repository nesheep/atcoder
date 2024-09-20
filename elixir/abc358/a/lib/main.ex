defmodule Main do
  def main do
    IO.read(:line)
    |> String.trim()
    |> case do
      "AtCoder Land" -> "Yes"
      _ -> "No"
    end
    |> IO.puts()
  end
end
