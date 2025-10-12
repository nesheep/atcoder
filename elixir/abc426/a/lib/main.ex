defmodule Main do
  @m %{
    "Ocelot" => 1,
    "Serval" => 2,
    "Lynx" => 3
  }

  def main do
    stdin_values()
    |> solve()
    |> output()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp solve([x, y]), do: Map.get(@m, x) >= Map.get(@m, y)

  defp output(true), do: "Yes"
  defp output(false), do: "No"
end
