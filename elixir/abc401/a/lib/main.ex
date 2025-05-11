defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> output()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input([s]), do: String.to_integer(s)

  defp output(s) when s >= 200 and s < 300, do: "Success"
  defp output(_), do: "Failure"
end
