//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sql

import (
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/uber/aresdb/common"
	queryCom "github.com/uber/aresdb/query/common"
	"github.com/uber/aresdb/query/sql/antlrgen"
	"github.com/uber/aresdb/query/sql/tree"
)

const (
	_aqlPrefix = "aql_"

	// supported query level
	maxlevelWith  = 1
	maxLevelQuery = 2

	// slice default size
	defaultSliceCap = 10

	// query types
	typeWithQuery = 1
	typeSubQuery  = 2
)

// ExprOrigin defines the expression origin
type ExprOrigin int

const (
	// ExprOriginWhere => the expression origin is from where clause
	ExprOriginWhere ExprOrigin = iota
	// ExprOriginJoinOn => the expression origin is from join on clause
	ExprOriginJoinOn
	// ExprOriginGroupBy => the expression origin is from groupingElement clause
	ExprOriginGroupBy
	// ExprOriginOthers => the expression origin is from other clauses case
	ExprOriginOthers
)

// SQL2AqlContext is the context of ASTVisitor
type SQL2AqlContext struct {
	/*
		Rules Of updating level, levelWith, levelQuery and mapXXX
		1. level: follow Treeprinter indent. Init value: 0
		2. levelWith: increase 1 if VisitWith is called. Init value: 0
		3. levelQuery: increase 1 if withQuery is called in VisitWith or VisitTableSubquery is called. Init value: 0
		4. mapXXX: create a new mapXXX[mapKey] if a new query is added (ie, VisitWithQuery or VisitTableSubquery). Init value: empty map table.
	*/
	// level is current tree level
	level int
	// levelWith is current with level
	levelWith int
	// levelQuery is current query level
	levelQuery int
	// MapQueryIdentifier is a mapping table. key=generateKey(...) value=arrayOfIdentifier.
	// Identifier can be namedQuery identifier or aliasedRelation identifier
	MapQueryIdentifier map[int]string
	// MapMeasures is a mapping table. key=generateKey(...) value=arrayOfMeasure
	MapMeasures map[int][]queryCom.Measure
	// MapDimensions is a mapping table. key=generateKey(...) value=arrayOfDimension
	MapDimensions map[int][]queryCom.Dimension
	// MapJoinTables is a mapping table. key=generateKey(...) value=arrayOfJoin
	MapJoinTables map[int][]queryCom.Join
	// MapRowFilters is a mapping table. key=generateKey(...) value=arrayOfRowFilter
	MapRowFilters map[int][]string
	// MapOrderBy is a mapping table. key=generateKey(...) value=arrayOfSortField
	MapOrderBy map[int][]queryCom.SortField
	// MapLimit is a mapping table. key=generateKey(...) value=arrayOfLimit
	MapLimit           map[int]int
	mapKey             int
	timeNow            int64
	timeFilter         queryCom.TimeFilter
	timezone           string
	exprOrigin         ExprOrigin
	fromJSON           []byte
	groupByJSON        []byte
	orderByJSON        []byte
	queryIdentifierSet map[string]int
	exprCheck          bool
	disableMainGroupBy bool
	exprLogicalOp      tree.LogicalBinaryExpType
}

// ASTBuilder is a visitor
type ASTBuilder struct {
	// Logger is a logger from appConfig
	Logger common.Logger
	// IStream is input antlr char stream
	IStream *antlr.CommonTokenStream
	// ParameterPosition is position in sql
	ParameterPosition int
	// SQL2AqlContext is the context of construncting AQL
	SQL2AqlCtx *SQL2AqlContext
	aql        *queryCom.AQLQuery

	// Flag that indicates whether aggregate function is seen
	aggFuncExists bool
}

func (v *ASTBuilder) defaultResult() interface{} { _ = "STUB: not implemented"; return nil }

func (v *ASTBuilder) shouldVisitNextChild(node antlr.RuleNode, currentResult interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *ASTBuilder) aggregateResult(node antlr.ParseTree, aggregate interface{}, nextResult interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (v *ASTBuilder) getQualifiedName(ctx antlrgen.IQualifiedNameContext) *tree.QualifiedName {
	_ = "STUB: not implemented"
	return nil
}

// VisitTerminal visits the node
func (v *ASTBuilder) VisitTerminal(node antlr.TerminalNode) interface{} {
	_ = "STUB: not implemented"

	// VisitErrorNode visits the node
	return nil
}

func (v *ASTBuilder) VisitErrorNode(node antlr.ErrorNode) interface{} {
	_ = "STUB: not implemented"

	// Visit visits the node
	return nil
}

func (v *ASTBuilder) Visit(tree antlr.ParseTree) interface{} { _ = "STUB: not implemented"; return nil }

// VisitChildren visits the node
func (v *ASTBuilder) VisitChildren(node antlr.RuleNode) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (v *ASTBuilder) visitIfPresent(ctx antlr.RuleContext, visitResult reflect.Type) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (v *ASTBuilder) visitList(ctxs []antlr.ParserRuleContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ********************** Visit SQL grammar starts ********************

// ********************** query expressions ********************

// VisitQuery visits the node
func (v *ASTBuilder) VisitQuery(ctx *antlrgen.QueryContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// handle with

// handle queryNoWith

// reset SQL2AqlContext

// VisitWith visits the node
func (v *ASTBuilder) VisitWith(ctx *antlrgen.WithContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitNamedQuery visits the node
func (v *ASTBuilder) VisitNamedQuery(ctx *antlrgen.NamedQueryContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// handle name

// handle columnAliases

// handle query

// VisitQueryNoWith visits the node
func (v *ASTBuilder) VisitQueryNoWith(ctx *antlrgen.QueryNoWithContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// handle queryTerm

// handle ORDER BY

// VisitQuerySpecification visits the node
func (v *ASTBuilder) VisitQuerySpecification(ctx *antlrgen.QuerySpecificationContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// handle from => join/table
// first process from clause so that subquery/withQuery identifier can be found in expression

// synthesize implicit join nodes

// handle select => measure

// handle subquery/withQuery with columnAliases,8
// subquery/withQuery columnalias has higher priority, ignore subquery/withQuery selectSingle identifier

// handle query or subquery/withQuery w/o columnAliases

// handle where => rowfilter/timefilter

// handle group by => dimension

// disable group by clause in manin query if with/subquery exists

// handle having => not support in AQL

// VisitSelectAll visits the node
func (v *ASTBuilder) VisitSelectAll(ctx *antlrgen.SelectAllContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitSelectSingle visits the node
func (v *ASTBuilder) VisitSelectSingle(ctx *antlrgen.SelectSingleContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitGroupBy visits the node
func (v *ASTBuilder) VisitGroupBy(ctx *antlrgen.GroupByContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitSingleGroupingSet visits the node
func (v *ASTBuilder) VisitSingleGroupingSet(ctx *antlrgen.SingleGroupingSetContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// timeBucket or numbericBucket is added into
// v.SQL2AqlCtx.MapDimensions[v.SQL2AqlCtx.mapKey] via visitFunctionCall

// VisitSortItem visits the node
func (v *ASTBuilder) VisitSortItem(ctx *antlrgen.SortItemContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ***************** boolean expressions ******************

// VisitExpression visits the node
func (v *ASTBuilder) VisitExpression(ctx *antlrgen.ExpressionContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitLogicalBinary visits the node
func (v *ASTBuilder) VisitLogicalBinary(ctx *antlrgen.LogicalBinaryContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitBooleanDefault visits the node
func (v *ASTBuilder) VisitBooleanDefault(ctx *antlrgen.BooleanDefaultContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitLogicalNot visits the node
func (v *ASTBuilder) VisitLogicalNot(ctx *antlrgen.LogicalNotContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// *************** from clause *****************

// VisitJoinRelation visits the node
func (v *ASTBuilder) VisitJoinRelation(ctx *antlrgen.JoinRelationContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitSampledRelation visits the node
func (v *ASTBuilder) VisitSampledRelation(ctx *antlrgen.SampledRelationContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitAliasedRelation visits the node
func (v *ASTBuilder) VisitAliasedRelation(ctx *antlrgen.AliasedRelationContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// handle identifier

// handle relationPrimary

// handle columnAliases

// VisitTableName visits the node
func (v *ASTBuilder) VisitTableName(ctx *antlrgen.TableNameContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// check if the table name is a withQ identifier

// VisitSubqueryRelation visits the node
func (v *ASTBuilder) VisitSubqueryRelation(ctx *antlrgen.SubqueryRelationContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// mapKey is the mapKey of parent query

// the index in v.SQL2AqlCtx.mapKey is the index of the aliasedRelation in parent from clause

// ********************* primary expressions **********************

// VisitFunctionCall visits the node
func (v *ASTBuilder) VisitFunctionCall(ctx *antlrgen.FunctionCallContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// 1. timebucket and numbericbucket are only from groupBy clause
// 2. timefilter is only from where clause

// VisitUnquotedIdentifier visits the node
func (v *ASTBuilder) VisitUnquotedIdentifier(ctx *antlrgen.UnquotedIdentifierContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitQuotedIdentifier visits the node
func (v *ASTBuilder) VisitQuotedIdentifier(ctx *antlrgen.QuotedIdentifierContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitDereference visits the node
func (v *ASTBuilder) VisitDereference(ctx *antlrgen.DereferenceContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// reject the expression if subquery/withQuery identifier is used in level 0 query

// ***************** Reserved *****************

// VisitStatementDefault visits the node
func (v *ASTBuilder) VisitStatementDefault(ctx *antlrgen.StatementDefaultContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitQueryTermDefault visits the node
func (v *ASTBuilder) VisitQueryTermDefault(ctx *antlrgen.QueryTermDefaultContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitSetOperation visits the node
func (v *ASTBuilder) VisitSetOperation(ctx *antlrgen.SetOperationContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitQueryPrimaryDefault visits the node
func (v *ASTBuilder) VisitQueryPrimaryDefault(ctx *antlrgen.QueryPrimaryDefaultContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitTable visits the node
func (v *ASTBuilder) VisitTable(ctx *antlrgen.TableContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitInlineTable visits the node
func (v *ASTBuilder) VisitInlineTable(ctx *antlrgen.InlineTableContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitSubquery visits the node
func (v *ASTBuilder) VisitSubquery(ctx *antlrgen.SubqueryContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitGroupingExpressions visits the node
func (v *ASTBuilder) VisitGroupingExpressions(ctx *antlrgen.GroupingExpressionsContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitSetQuantifier visits the node
func (v *ASTBuilder) VisitSetQuantifier(ctx *antlrgen.SetQuantifierContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitRelationDefault visits the node
func (v *ASTBuilder) VisitRelationDefault(ctx *antlrgen.RelationDefaultContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitJoinType visits the node
func (v *ASTBuilder) VisitJoinType(ctx *antlrgen.JoinTypeContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitJoinCriteria visits the node
func (v *ASTBuilder) VisitJoinCriteria(ctx *antlrgen.JoinCriteriaContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitSampleType visits the node
func (v *ASTBuilder) VisitSampleType(ctx *antlrgen.SampleTypeContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitColumnAliases visits the node
func (v *ASTBuilder) VisitColumnAliases(ctx *antlrgen.ColumnAliasesContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitParenthesizedRelation visits the node
func (v *ASTBuilder) VisitParenthesizedRelation(ctx *antlrgen.ParenthesizedRelationContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitQuantifiedComparison visits the node
func (v *ASTBuilder) VisitQuantifiedComparison(ctx *antlrgen.QuantifiedComparisonContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitBetween visits the node
func (v *ASTBuilder) VisitBetween(ctx *antlrgen.BetweenContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitInList visits the node
func (v *ASTBuilder) VisitInList(ctx *antlrgen.InListContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitInSubquery visits the node
func (v *ASTBuilder) VisitInSubquery(ctx *antlrgen.InSubqueryContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitValueExpressionDefault visits the node
func (v *ASTBuilder) VisitValueExpressionDefault(ctx *antlrgen.ValueExpressionDefaultContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitConcatenation visits the node
func (v *ASTBuilder) VisitConcatenation(ctx *antlrgen.ConcatenationContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitArithmeticBinary visits the node
func (v *ASTBuilder) VisitArithmeticBinary(ctx *antlrgen.ArithmeticBinaryContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitArithmeticUnary visits the node
func (v *ASTBuilder) VisitArithmeticUnary(ctx *antlrgen.ArithmeticUnaryContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitAtTimeZone visits the node
func (v *ASTBuilder) VisitAtTimeZone(ctx *antlrgen.AtTimeZoneContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitTypeConstructor visits the node
func (v *ASTBuilder) VisitTypeConstructor(ctx *antlrgen.TypeConstructorContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitSpecialDateTimeFunction visits the node
func (v *ASTBuilder) VisitSpecialDateTimeFunction(ctx *antlrgen.SpecialDateTimeFunctionContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitPredicated visits the node
func (v *ASTBuilder) VisitPredicated(ctx *antlrgen.PredicatedContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitComparison visits the node
func (v *ASTBuilder) VisitComparison(ctx *antlrgen.ComparisonContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitParenthesizedExpression visits the node
func (v *ASTBuilder) VisitParenthesizedExpression(ctx *antlrgen.ParenthesizedExpressionContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitIntervalLiteral visits the node
func (v *ASTBuilder) VisitIntervalLiteral(ctx *antlrgen.IntervalLiteralContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitNumericLiteral visits the node
func (v *ASTBuilder) VisitNumericLiteral(ctx *antlrgen.NumericLiteralContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitBooleanLiteral visits the node
func (v *ASTBuilder) VisitBooleanLiteral(ctx *antlrgen.BooleanLiteralContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitColumnReference \visits the node
func (v *ASTBuilder) VisitColumnReference(ctx *antlrgen.ColumnReferenceContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitNullLiteral visits the node
func (v *ASTBuilder) VisitNullLiteral(ctx *antlrgen.NullLiteralContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitRowConstructor visits the node
func (v *ASTBuilder) VisitRowConstructor(ctx *antlrgen.RowConstructorContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitSubscript visits the node
func (v *ASTBuilder) VisitSubscript(ctx *antlrgen.SubscriptContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitSubqueryExpression visits the node
func (v *ASTBuilder) VisitSubqueryExpression(ctx *antlrgen.SubqueryExpressionContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitBinaryLiteral visits the node
func (v *ASTBuilder) VisitBinaryLiteral(ctx *antlrgen.BinaryLiteralContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitCurrentUser visits the node
func (v *ASTBuilder) VisitCurrentUser(ctx *antlrgen.CurrentUserContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitStringLiteral visits the node
func (v *ASTBuilder) VisitStringLiteral(ctx *antlrgen.StringLiteralContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitArrayConstructor visits the node
func (v *ASTBuilder) VisitArrayConstructor(ctx *antlrgen.ArrayConstructorContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitGroupingOperation visits the node
func (v *ASTBuilder) VisitGroupingOperation(ctx *antlrgen.GroupingOperationContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitBasicStringLiteral visits the node
func (v *ASTBuilder) VisitBasicStringLiteral(ctx *antlrgen.BasicStringLiteralContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitUnicodeStringLiteral visits the node
func (v *ASTBuilder) VisitUnicodeStringLiteral(ctx *antlrgen.UnicodeStringLiteralContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitTimeZoneInterval visits the node
func (v *ASTBuilder) VisitTimeZoneInterval(ctx *antlrgen.TimeZoneIntervalContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitTimeZoneString visits the node
func (v *ASTBuilder) VisitTimeZoneString(ctx *antlrgen.TimeZoneStringContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitComparisonOperator visits the node
func (v *ASTBuilder) VisitComparisonOperator(ctx *antlrgen.ComparisonOperatorContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitComparisonQuantifier visits the node
func (v *ASTBuilder) VisitComparisonQuantifier(ctx *antlrgen.ComparisonQuantifierContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitBooleanValue visits the node
func (v *ASTBuilder) VisitBooleanValue(ctx *antlrgen.BooleanValueContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitInterval visits the node
func (v *ASTBuilder) VisitInterval(ctx *antlrgen.IntervalContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitIntervalField visits the node
func (v *ASTBuilder) VisitIntervalField(ctx *antlrgen.IntervalFieldContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitNormalForm visits the node
func (v *ASTBuilder) VisitNormalForm(ctx *antlrgen.NormalFormContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitSqltype visits the node
func (v *ASTBuilder) VisitSqltype(ctx *antlrgen.SqltypeContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitTypeParameter visits the node
func (v *ASTBuilder) VisitTypeParameter(ctx *antlrgen.TypeParameterContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitBaseType visits the node
func (v *ASTBuilder) VisitBaseType(ctx *antlrgen.BaseTypeContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitWhenClause visits the node
func (v *ASTBuilder) VisitWhenClause(ctx *antlrgen.WhenClauseContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitFilter visits the node
func (v *ASTBuilder) VisitFilter(ctx *antlrgen.FilterContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitQualifiedName visits the node
func (v *ASTBuilder) VisitQualifiedName(ctx *antlrgen.QualifiedNameContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitBackQuotedIdentifier visits the node
func (v *ASTBuilder) VisitBackQuotedIdentifier(ctx *antlrgen.BackQuotedIdentifierContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitDigitIdentifier visits the node
func (v *ASTBuilder) VisitDigitIdentifier(ctx *antlrgen.DigitIdentifierContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitDecimalLiteral visits the node
func (v *ASTBuilder) VisitDecimalLiteral(ctx *antlrgen.DecimalLiteralContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitDoubleLiteral visits the node
func (v *ASTBuilder) VisitDoubleLiteral(ctx *antlrgen.DoubleLiteralContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitIntegerLiteral visits the node
func (v *ASTBuilder) VisitIntegerLiteral(ctx *antlrgen.IntegerLiteralContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// VisitNonReserved visits the node
func (v *ASTBuilder) VisitNonReserved(ctx *antlrgen.NonReservedContext) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ***************** helpers *****************
type orderByContext interface {
	ORDER() antlr.TerminalNode
	AllSortItem() []antlrgen.ISortItemContext
}

func (v *ASTBuilder) getOrderBy(ctx orderByContext) *tree.OrderBy {
	_ = "STUB: not implemented"
	return nil
}

func (v *ASTBuilder) getLocation(input interface{}) *tree.NodeLocation {
	_ = "STUB: not implemented"
	return nil
}

// getText extracts string from original input sql
func (v *ASTBuilder) getText(ctx antlr.ParserRuleContext) string {
	_ = "STUB: not implemented"
	return ""
}

func (v *ASTBuilder) setTimefilter(ctx []antlrgen.IExpressionContext) {
	_ = "STUB: not implemented"
	return
}

func (v *ASTBuilder) setTimeNow(ctx []antlrgen.IExpressionContext) {
	_ = "STUB: not implemented"
	return
}

func (v *ASTBuilder) setNumericBucketizer(ctx []antlrgen.IExpressionContext, def string) {
	_ = "STUB: not implemented"
	return
}

// mergeWithOrSubQueries merge all subquery/withQuery's information into v.aql
func (v *ASTBuilder) mergeWithOrSubQueries() { _ = "STUB: not implemented"; return }

// mergeWithOrSubQuery merge one subquery/withQuery information into v.aql
func (v *ASTBuilder) mergeWithOrSubQuery(key int, ignoreJoin bool) {
	_ = "STUB: not implemented"
	return
}

// case1: this measure of subquery/withQuery is not a supportingMeasure

// case2: this measure of subquery/withQuery is a supportingMeasure

// isMeasureInMain check a measure of subquery/withQuery is also a measure of level 0 query
// return the index of the measure in level 0 query; otherwise return -1
func (v *ASTBuilder) isMeasureInMain(key, index int) int { _ = "STUB: not implemented"; return 0 }

// GetAQL construct AQLQuery via read through SQL2AqlCtx
func (v *ASTBuilder) GetAQL() *queryCom.AQLQuery { _ = "STUB: not implemented"; return nil }

// there is no subquery/withQuery in sql

// remove measures that should be dimensions

// GetTextIfPresent visits the node
func (v *ASTBuilder) GetTextIfPresent(token antlr.Token) string {
	_ = "STUB: not implemented"
	return ""
}

// isDistinct check if DISTINCT quantifier is set
func (v *ASTBuilder) isDistinct(setQuantifier antlrgen.ISetQuantifierContext) bool {
	_ = "STUB: not implemented"
	return false
}

// getLogicalBinaryOperator returns an input token's logicalBinaryExpression operator type
func (v *ASTBuilder) getLogicalBinaryOperator(token int) tree.LogicalBinaryExpType {
	_ = "STUB: not implemented"
	return *new(tree.LogicalBinaryExpType)
}

func (v *ASTBuilder) getJoinType(ctx *antlrgen.JoinRelationContext) tree.JoinType {
	_ = "STUB: not implemented"
	return *new(tree.JoinType)
}

func (v *ASTBuilder) getCtxLevels(s2aCtx *SQL2AqlContext) (level, levelWith, levelQuery int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func (v *ASTBuilder) setCtxLevels(s2aCtx *SQL2AqlContext, level, levelWith, levelQuery int) {
	_ = "STUB: not implemented"
	return
}

// generateKey constructs mapKey based on levelQuery and index of the query at the current levelQuery
func (v *ASTBuilder) generateKey(qLevel, qType, index int) int { _ = "STUB: not implemented"; return 0 }

func (v *ASTBuilder) getInfoByKey(mapKey int) (qLevel, qType, index int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func (v *ASTBuilder) isValidWithOrSubQuery(s2aCtx *SQL2AqlContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// check if from clause in main query(ie. qLevel = 0) mix table with subquery/withQuery

// check if all subquery/withQuery from clauses are same

// exit if no subquery/withQuery

// subquery has no identifier

// AQL requires that the first level query is either from tables or from subqueries/withQuery
func (v *ASTBuilder) isQueryFromMixed(s2aCtx *SQL2AqlContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AQL requires that all subqueries or withQuery from clauses are same
func (v *ASTBuilder) isSameFromTables(s2aCtx *SQL2AqlContext, mapKey int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// generte from clause json bytes based on the first subquery/withQuery

// compare current from clause with ctx.fromJSON

// AQL requires that all subqueries or withQuery groupBy clauses are sameo
func (v *ASTBuilder) isSameGroupBy(s2aCtx *SQL2AqlContext, mapKey int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// generte group by clause json bytes based on the first subquery/withQuery

// compare current groupBy clause with ctx.groupByJSON

// AQL requires that all subqueries or withQuery orderBy clauses are same
func (v *ASTBuilder) isSameOrderBy(s2aCtx *SQL2AqlContext, mapKey int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// generte group by clause json bytes based on the first subquery/withQuery

// compare current groupBy clause with ctx.groupByJSON

// addQIdentifier adds subquery/withQuery identifier and its mapKey into queryIdentifierSet
func (v *ASTBuilder) addQIdentifier(s2aCtx *SQL2AqlContext, indentifier string, key int) error {
	_ = "STUB: not implemented"
	return nil
}

// isWithQueryIdentifier check if name is a withQuery identifier
func (v *ASTBuilder) isWithQueryIdentifier(s2aCtx *SQL2AqlContext, name string) bool {
	_ = "STUB: not implemented"
	return false
}

// isSubOrWithQueryIdentifier check if name is a subquery/withQuery identifier
func (v *ASTBuilder) isSubOrWithQueryIdentifier(s2aCtx *SQL2AqlContext, name string) int {
	_ = "STUB: not implemented"
	return 0
}

// lookupSQLExpr is used by groupBy, orderBy, having clause whose sql expression is select column alias.
// It returns the select column alias and sql expresion.
func (v *ASTBuilder) lookupSQLExpr(s2aCtx *SQL2AqlContext, mapKey int, str string) (alias, sqlExpr string) {
	_ = "STUB: not implemented"
	return "", ""
}

// check path from expression node to leaf node has logicalBinaryOperator OR
func (v *ASTBuilder) hasORInPath(node *antlr.BaseParserRuleContext) tree.LogicalBinaryExpType {
	_ = "STUB: not implemented"
	return *new(tree.LogicalBinaryExpType)
}

// Parse parses input sql
func Parse(sql string, logger common.Logger) (aql *queryCom.AQLQuery, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Setup the input sql

// Create the Lexer

// Create the Parser

// Finally parse the sql

// Construct ASTBuilder

// non agg query overwrite
